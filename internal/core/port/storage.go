package port

import "github.com/marius-go/computer-management-service/internal/core/domain"

type ComputerStorage interface {
	CreateComputer(domain.Computer) (domain.Computer, error)
	UpdateComputer(domain.Computer) (domain.Computer, error)
	DeleteComputer(name string) error
	GetComputer(name string) (domain.Computer, error)
	ListComputers(employeeAbbreviation *string) ([]domain.Computer, error)
}
