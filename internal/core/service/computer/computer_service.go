package computer

import (
	"github.com/marius-go/computer-management-service/internal/core/domain"
	"github.com/marius-go/computer-management-service/internal/core/port"
)

type Service struct {
	storage port.ComputerStorage
}

func New(storage port.ComputerStorage) *Service {
	return &Service{storage: storage}
}

func (s *Service) CreateComputer(domain.Computer) (domain.Computer, error) {
	return domain.Computer{}, nil
}

func (s *Service) UpdateComputer(domain.Computer) (domain.Computer, error) {
	return domain.Computer{}, nil
}

func (s *Service) DeleteComputer(name string) error {
	return nil
}

func (s *Service) GetComputer(name string) (domain.Computer, error) {
	return domain.Computer{}, nil
}

func (s *Service) ListComputers(employeeAbbreviation string) ([]domain.Computer, error) {
	return nil, nil
}
