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

func (s *MemoryStorage) CreateComputer(domain.Computer) (domain.Computer, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return domain.Computer{}, nil
}

func (s *MemoryStorage) UpdateComputer(domain.Computer) (domain.Computer, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return domain.Computer{}, nil
}

func (s *MemoryStorage) DeleteComputer(name string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return nil
}

func (s *MemoryStorage) GetComputer(name string) (domain.Computer, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return domain.Computer{}, nil
}

func (s *MemoryStorage) ListComputers(employeeAbbreviation string) ([]domain.Computer, error) {
	s.lock.Lock()
	defer s.lock.RUnlock()
	return nil, nil
}
