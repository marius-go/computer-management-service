package memstorage

import (
	"sync"

	"github.com/marius-go/computer-management-service/internal/core/domain"
)

type MemoryStorage struct {
	computers map[string]computer
	lock      sync.RWMutex
}

func New() *MemoryStorage {
	computers := make(map[string]computer)
	return &MemoryStorage{computers: computers}
}

func (s *MemoryStorage) CreateComputer(computer domain.Computer) (domain.Computer, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	comp := ComputerFromDomain(computer)
	s.computers[computer.Name] = comp

	return comp.ToDomain(), nil
}

func (s *MemoryStorage) UpdateComputer(computer domain.Computer) (domain.Computer, error) {
	return s.CreateComputer(computer)
}

func (s *MemoryStorage) DeleteComputer(name string) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.computers, name)

	return nil
}

func (s *MemoryStorage) GetComputer(name string) (domain.Computer, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	computer, ok := s.computers[name]
	if !ok {
		return domain.Computer{}, domain.ErrNotFound
	}

	return computer.ToDomain(), nil
}

func (s *MemoryStorage) ListComputers(employeeAbbreviation *string) ([]domain.Computer, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	computers := []domain.Computer{}

	for _, computer := range s.computers {
		if employeeAbbreviation == nil || computer.EmployeeAbbreviation == *employeeAbbreviation {
			computers = append(computers, computer.ToDomain())
		}
	}

	return computers, nil
}
