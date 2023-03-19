package memstorage

import "github.com/marius-go/computer-management-service/internal/core/domain"

type computer struct {
	Name                 string
	MAC                  *string
	IP                   *string
	EmployeeAbbreviation *string
	Description          *string
}

func (c computer) ToDomain() domain.Computer {
	return domain.Computer(c)
}

func ComputerFromDomain(comp domain.Computer) computer {
	return computer(comp)
}
