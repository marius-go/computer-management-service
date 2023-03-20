package computer

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"sort"
	"sync"
	"testing"

	"github.com/marius-go/computer-management-service/internal/core/domain"
	"github.com/marius-go/computer-management-service/internal/infrastructure/memstorage"
)

type MockAdminNotifier struct {
	notificationCount int
	lock              sync.RWMutex
}

func (n *MockAdminNotifier) Notify(level string, employeeAbbreviation string, message string) error {
	n.lock.Lock()
	defer n.lock.Unlock()

	n.notificationCount++
	return nil
}

func (n *MockAdminNotifier) getNotificationCount() int {
	n.lock.RLock()
	defer n.lock.RUnlock()

	return n.notificationCount
}

func toPtr[T any](value T) *T {
	return &value
}

func initializeComputerService(initialComputers ...domain.Computer) (*Service, error) {
	storage := memstorage.New()
	mockNotifier := MockAdminNotifier{}
	computerService := New(storage, &mockNotifier)

	for _, comp := range initialComputers {
		_, err := computerService.CreateComputer(comp)
		if err != nil {
			return nil, fmt.Errorf("could not add computer: %w", err)
		}
	}
	return computerService, nil
}

// Note: the service always returns computer with all properties initialized,
// so for proper comparison we need to initialize all properties
func initializeOptionalProperties(computer *domain.Computer) {
	if computer.EmployeeAbbreviation == nil {
		computer.EmployeeAbbreviation = toPtr("")
	}
	if computer.Description == nil {
		computer.Description = toPtr("")
	}
}

var validComputer = domain.Computer{Name: "existing", MAC: toPtr("testMAC"), IP: toPtr("0.0.0.0")}

func TestService_CreateComputer(t *testing.T) {
	initialComputer := validComputer

	tests := []struct {
		name              string
		computer          domain.Computer
		wantErr           bool
		wantErrConflict   bool
		wantErrValidation bool
	}{
		{
			name:     "creating computer with all properties set succeeds",
			computer: domain.Computer{Name: "testPc", MAC: toPtr("testMAC"), IP: toPtr("0.0.0.0"), EmployeeAbbreviation: toPtr("mmu"), Description: toPtr("test computer")},
			wantErr:  false,
		},
		{
			name:     "creating computer with only required properties succeeds",
			computer: domain.Computer{Name: "testPc", MAC: toPtr("testMAC"), IP: toPtr("0.0.0.0")},
			wantErr:  false,
		},
		{
			name:              "creating computer with missing required values fails",
			computer:          domain.Computer{Name: "testPc"},
			wantErr:           true,
			wantErrValidation: true,
		},
		{
			name:              "creating computer with invalid employee abbreviation fails",
			computer:          domain.Computer{Name: "testPc", MAC: toPtr("testMAC"), IP: toPtr("0.0.0.0"), EmployeeAbbreviation: toPtr("too long")},
			wantErr:           true,
			wantErrValidation: true,
		},
		{
			name:            "creating computer with already existing name fails",
			computer:        initialComputer,
			wantErr:         true,
			wantErrConflict: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockNotifier := MockAdminNotifier{}
			storage := memstorage.New()
			storage.CreateComputer(initialComputer)
			computerService := New(storage, &mockNotifier)

			gotComputer, err := computerService.CreateComputer(tt.computer)
			if err != nil {
				if !tt.wantErr {
					t.Error("did not want error, got:", err)
				}

				var errConflict domain.ErrConflict
				if tt.wantErrConflict && !errors.As(err, &errConflict) {
					t.Error("want conflict error, got:", err)
				}

				var errValidation domain.ErrValidation
				if tt.wantErrValidation && !errors.As(err, &errValidation) {
					t.Error("want validation error, got:", err)
				}

				return // nothing more to test if an error occurred
			}

			wantComputer := tt.computer
			initializeOptionalProperties(&wantComputer)

			if !reflect.DeepEqual(gotComputer, wantComputer) {
				t.Error("computer in response has not correct values, want:", wantComputer, "got:", gotComputer)
				return
			}

			// check if the create response matches the get result
			gotComputer, err = computerService.GetComputer(tt.computer.Name)
			if err != nil {
				t.Error("could not get computer:", err)
				return
			}
			if !reflect.DeepEqual(gotComputer, wantComputer) {
				t.Error("added computer has not correct values, want:", wantComputer, "got:", gotComputer)
				return
			}
		})
	}
}

func TestService_GetComputer(t *testing.T) {
	initialComputer := validComputer

	tests := []struct {
		name         string
		computerName string
		wantErr      error
	}{
		{
			name:         "get computer which was inserted before",
			computerName: initialComputer.Name,
			wantErr:      nil,
		},
		{
			name:         "not found error if computer does not exist",
			computerName: "non-existent",
			wantErr:      domain.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computerService, err := initializeComputerService(initialComputer)
			if err != nil {
				t.Error("could not initialize computer service:", err)
			}

			gotComputer, err := computerService.GetComputer(tt.computerName)
			if err != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Error("wantError:", tt.wantErr, "gotError:", err)
					return
				}
				return // nothing more to test if an error occurred
			}

			wantComputer := initialComputer
			initializeOptionalProperties(&wantComputer)

			if !reflect.DeepEqual(gotComputer, wantComputer) {
				t.Error("added computer has not correct values, want:", wantComputer, "got:", gotComputer)
				return
			}
		})
	}
}

func TestService_UpdateComputer(t *testing.T) {
	initialComputer := validComputer

	tests := []struct {
		name              string
		update            domain.Computer
		wantComputer      domain.Computer
		wantErr           bool
		wantErrNotFound   bool
		wantErrValidation bool
	}{
		{
			name:   "only set properties are updated",
			update: domain.Computer{Name: initialComputer.Name, MAC: toPtr("new MAC")},
			wantComputer: func() domain.Computer {
				wantComputer := initialComputer
				*wantComputer.MAC = "new MAC"
				return wantComputer
			}(),
			wantErr: false,
		},
		{
			name:            "fails if name does not exist",
			update:          domain.Computer{Name: "non-existent"},
			wantErr:         true,
			wantErrNotFound: true,
		},
		{
			name:              "fails if invalid employee abbreviation is set",
			update:            domain.Computer{Name: initialComputer.Name, EmployeeAbbreviation: toPtr("too long")},
			wantErr:           true,
			wantErrValidation: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computerService, err := initializeComputerService(initialComputer)
			if err != nil {
				t.Error("could not initialize computer service:", err)
			}

			gotComputer, err := computerService.UpdateComputer(tt.update)
			if err != nil {
				if !tt.wantErr {
					t.Error("did not want error, got:", err)
				}

				if tt.wantErrNotFound && !errors.Is(err, domain.ErrNotFound) {
					t.Error("wantError:", domain.ErrNotFound, "gotError:", err)
				}

				var errValidation domain.ErrValidation
				if tt.wantErrValidation && !errors.As(err, &errValidation) {
					t.Error("want validation error, got:", err)
				}

				return // nothing more to test if an error occurred
			}

			initializeOptionalProperties(&tt.wantComputer)
			if !reflect.DeepEqual(gotComputer, tt.wantComputer) {
				t.Error("computer in response has not correct values, want:", tt.wantComputer, "got:", gotComputer)
				return
			}

			// check if the update response matches the get result
			gotComputer, err = computerService.GetComputer(tt.update.Name)
			if err != nil {
				t.Error("could not get computer:", err)
			}
			if !reflect.DeepEqual(gotComputer, tt.wantComputer) {
				t.Error("updated computer has not correct values, want:", tt.wantComputer, "got:", gotComputer)
				return
			}
		})
	}
}

func TestService_DeleteComputer(t *testing.T) {
	initialComputer := validComputer

	tests := []struct {
		name         string
		computerName string
		wantErr      error
	}{
		{
			name:         "existing computer is deleted",
			computerName: initialComputer.Name,
			wantErr:      nil,
		},
		{
			name:         "fails if computer does not exist",
			computerName: "non-existent",
			wantErr:      domain.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computerService, err := initializeComputerService(initialComputer)
			if err != nil {
				t.Error("could not initialize computer service:", err)
			}

			err = computerService.DeleteComputer(tt.computerName)
			if err != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Error("wantError:", tt.wantErr, "gotError:", err)
					return
				}

				return // nothing more to test if an error occurred
			}

			// check if removed computer is actually gone
			computer, err := computerService.GetComputer(tt.computerName)
			if !errors.Is(err, domain.ErrNotFound) {
				t.Error("deleted computer still exists", computer)
				return
			}
		})
	}
}

func TestService_AdminNotification(t *testing.T) {
	type op int
	const (
		create op = iota
		update
	)

	initialComputers := []domain.Computer{
		{Name: "pc1", MAC: toPtr("mac1"), IP: toPtr("ip1")},
		{Name: "pc2", MAC: toPtr("mac2"), IP: toPtr("ip2")},
		{Name: "pc3", MAC: toPtr("mac3"), IP: toPtr("ip3"), EmployeeAbbreviation: toPtr("mmu")},
		{Name: "pc4", MAC: toPtr("mac4"), IP: toPtr("ip4"), EmployeeAbbreviation: toPtr("mmu")},
	}

	tests := []struct {
		name                   string
		computersToAddOrUpdate []domain.Computer
		wantNotifications      int
		operation              op
	}{
		{
			name: "two notifications for 3 and 4 assigned computer via create",
			computersToAddOrUpdate: []domain.Computer{
				{Name: "pc5", MAC: toPtr("mac5"), IP: toPtr("ip5"), EmployeeAbbreviation: toPtr("mmu")},
				{Name: "pc6", MAC: toPtr("mac6"), IP: toPtr("ip6"), EmployeeAbbreviation: toPtr("mmu")},
			},
			operation:         create,
			wantNotifications: 2,
		},
		{
			name: "two notifications for 3 and 4 assigned computer via update",
			computersToAddOrUpdate: []domain.Computer{
				{Name: "pc1", EmployeeAbbreviation: toPtr("mmu")},
				{Name: "pc2", EmployeeAbbreviation: toPtr("mmu")},
			},
			operation:         update,
			wantNotifications: 2,
		},
		{
			name: "no additional notification when removing assignment, even if employee has still 3 computers assigned",
			computersToAddOrUpdate: []domain.Computer{
				{Name: "pc1", EmployeeAbbreviation: toPtr("mmu")},
				{Name: "pc2", EmployeeAbbreviation: toPtr("mmu")},
				{Name: "pc2", EmployeeAbbreviation: toPtr("emu")},
			},
			operation:         update,
			wantNotifications: 2,
		},
		{
			name: "no notifications < 3 assigned computer via create",
			computersToAddOrUpdate: []domain.Computer{
				{Name: "pc5", MAC: toPtr("mac5"), IP: toPtr("ip5"), EmployeeAbbreviation: toPtr("emu")},
			},
			operation:         create,
			wantNotifications: 0,
		},
		{
			name: "no notifications < 3 assigned computer via update",
			computersToAddOrUpdate: []domain.Computer{
				{Name: "pc1", EmployeeAbbreviation: toPtr("emu")},
			},
			operation:         update,
			wantNotifications: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storage := memstorage.New()
			mockNotifier := MockAdminNotifier{}
			computerService := New(storage, &mockNotifier)

			for _, comp := range initialComputers {
				_, err := computerService.CreateComputer(comp)
				if err != nil {
					t.Error("could not add computer:", err)
					return
				}
			}

			switch tt.operation {
			case create:
				for _, comp := range tt.computersToAddOrUpdate {
					_, err := computerService.CreateComputer(comp)
					if err != nil {
						t.Error("could not add computer:", err)
						return
					}
					list, _ := computerService.ListComputers(toPtr("mmu"))
					log.Println(len(list))
				}
			case update:
				for _, comp := range tt.computersToAddOrUpdate {
					_, err := computerService.UpdateComputer(comp)
					if err != nil {
						t.Error("could not add computer:", err)
						return
					}
				}
			}

			gotNotifications := mockNotifier.getNotificationCount()
			if gotNotifications != tt.wantNotifications {
				t.Error("did not get the correct count of notifications, want:", tt.wantNotifications, "got:", gotNotifications)
			}
		})
	}
}

func TestService_ListComputer(t *testing.T) {
	ownerless1 := domain.Computer{Name: "pc1", MAC: toPtr("mac1"), IP: toPtr("ip1")}
	ownerless2 := domain.Computer{Name: "pc2", MAC: toPtr("mac2"), IP: toPtr("ip2")}
	mmu1 := domain.Computer{Name: "pc3", MAC: toPtr("mac3"), IP: toPtr("ip3"), EmployeeAbbreviation: toPtr("mmu")}
	mmu2 := domain.Computer{Name: "pc4", MAC: toPtr("mac4"), IP: toPtr("ip4"), EmployeeAbbreviation: toPtr("mmu")}
	emu1 := domain.Computer{Name: "pc5", MAC: toPtr("mac5"), IP: toPtr("ip5"), EmployeeAbbreviation: toPtr("emu")}

	initialComputers := []domain.Computer{
		ownerless1,
		ownerless2,
		mmu1,
		mmu2,
		emu1,
	}

	tests := []struct {
		name                 string
		employeeAbbreviation *string
		wantComputers        []domain.Computer
	}{
		{
			name:                 "get all computers if no filter is set",
			employeeAbbreviation: nil,
			wantComputers:        []domain.Computer{ownerless1, ownerless2, mmu1, mmu2, emu1},
		},
		{
			name:                 "get all ownerless computers if empty employee abbreviation is set",
			employeeAbbreviation: toPtr(""),
			wantComputers:        []domain.Computer{ownerless1, ownerless2},
		},
		{
			name:                 "get all computers of an employee abbreviation",
			employeeAbbreviation: toPtr("mmu"),
			wantComputers:        []domain.Computer{mmu1, mmu2},
		},
		{
			name:                 "get empty list if employee abbreviation does not exist",
			employeeAbbreviation: toPtr("non-existent"),
			wantComputers:        []domain.Computer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			computerService, err := initializeComputerService(initialComputers...)
			if err != nil {
				t.Error("could not initialize computer service:", err)
			}

			gotComputers, err := computerService.ListComputers(tt.employeeAbbreviation)
			if err != nil {
				t.Error("ListComputers failed:", err)
				return
			}

			sort.Slice(gotComputers, func(i, j int) bool { return gotComputers[i].Name < gotComputers[j].Name })
			sort.Slice(tt.wantComputers, func(i, j int) bool { return tt.wantComputers[i].Name < tt.wantComputers[j].Name })
			for ii := 0; ii < len(tt.wantComputers); ii++ {
				initializeOptionalProperties(&tt.wantComputers[ii])
			}

			if !reflect.DeepEqual(gotComputers, tt.wantComputers) {
				t.Error("did not get expected list of computers, want:", tt.wantComputers, "got:", gotComputers)
				return
			}
		})
	}
}
