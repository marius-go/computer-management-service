package computer

import (
	"errors"
	"fmt"
	"log"

	"github.com/marius-go/computer-management-service/internal/core/domain"
	"github.com/marius-go/computer-management-service/internal/core/port"
)

const assignedComputersThreshold int = 3

type Service struct {
	storage       port.ComputerStorage
	adminNotifier port.AdminNotification
}

func New(storage port.ComputerStorage, adminNotifier port.AdminNotification) *Service {
	return &Service{storage: storage, adminNotifier: adminNotifier}
}

func (s *Service) CreateComputer(computer domain.Computer) (domain.Computer, error) {
	err := computer.Validate()
	if err != nil {
		return domain.Computer{}, err
	}

	_, err = s.storage.GetComputer(computer.Name)
	// only continue if resource is not found
	if err == nil {
		return domain.Computer{}, domain.ErrConflict("Name")
	} else if !errors.Is(err, domain.ErrNotFound) {
		return domain.Computer{}, fmt.Errorf("storage error: %w", err)
	}

	comp, err := s.storage.CreateComputer(computer)
	if err != nil {
		return domain.Computer{}, annotateStorageError(err)
	}

	// check for assignments only after the computer is added, due to races we could get to many notifications,
	// but when checking before adding, notifications events could be missed due to those races
	if computer.EmployeeAbbreviation != nil && *computer.EmployeeAbbreviation != "" {
		s.checkAssignments(*computer.EmployeeAbbreviation)
	}

	return comp, nil
}

func (s *Service) UpdateComputer(update domain.Computer) (domain.Computer, error) {

	computer, err := s.storage.GetComputer(update.Name)
	if err != nil {
		return domain.Computer{}, annotateStorageError(err)
	}

	updateComputer(&computer, update)

	err = computer.Validate()
	if err != nil {
		return domain.Computer{}, err
	}

	computer, err = s.storage.UpdateComputer(computer)
	if err != nil {
		return domain.Computer{}, annotateStorageError(err)
	}

	// check for assignments only after the computer was updated, due to races we could get to many notifications,
	// but when checking before updated, notifications events could be missed due to those races
	if update.EmployeeAbbreviation != nil && *update.EmployeeAbbreviation != "" {
		s.checkAssignments(*update.EmployeeAbbreviation)
	}

	return computer, nil
}

func (s *Service) DeleteComputer(name string) error {
	_, err := s.storage.GetComputer(name)
	if err != nil {
		return annotateStorageError(err)
	}

	err = s.storage.DeleteComputer(name)
	if err != nil {
		return annotateStorageError(err)
	}

	return nil
}

func (s *Service) GetComputer(name string) (domain.Computer, error) {
	computer, err := s.storage.GetComputer(name)
	if err != nil {
		return domain.Computer{}, annotateStorageError(err)
	}

	return computer, nil
}

func (s *Service) ListComputers(employeeAbbreviation *string) ([]domain.Computer, error) {
	computers, err := s.storage.ListComputers(employeeAbbreviation)
	if err != nil {
		return nil, annotateStorageError(err)
	}

	return computers, nil
}

func (s *Service) checkAssignments(employeeAbbreviation string) {
	assignedComputers, err := s.storage.ListComputers(&employeeAbbreviation)
	if err != nil {
		log.Println("could not check assignments due to storage error:", err)
		return
	}

	numAssignedComputers := len(assignedComputers)
	if numAssignedComputers >= assignedComputersThreshold {
		message := fmt.Sprint("employee has ", numAssignedComputers, " computers assigned")
		err := s.adminNotifier.Notify("warning", employeeAbbreviation, message)
		if err != nil {
			log.Println("could not notify due to notification service error:", err, "; dropped notification:", message, "("+employeeAbbreviation+")")
		}
	}
}

// updates `computer` with the set values of `update`
func updateComputer(computer *domain.Computer, update domain.Computer) {
	if update.MAC != nil {
		*computer.MAC = *update.MAC
	}
	if update.IP != nil {
		*computer.IP = *update.IP
	}
	if update.EmployeeAbbreviation != nil {
		*computer.EmployeeAbbreviation = *update.EmployeeAbbreviation
	}
	if update.Description != nil {
		*computer.Description = *update.Description
	}
}

func annotateStorageError(err error) error {
	switch {
	case errors.Is(err, domain.ErrNotFound):
		return err
	default:
		return fmt.Errorf("storage error: %w", err)
	}
}
