package memstorage

import "github.com/marius-go/computer-management-service/internal/core/domain"

type computer struct {
	Name                 string
	MAC                  string
	IP                   string
	EmployeeAbbreviation string
	Description          string
}

func (c computer) ToDomain() domain.Computer {
	return domain.Computer{
		Name:                 c.Name,
		MAC:                  &c.MAC,
		IP:                   &c.IP,
		EmployeeAbbreviation: &c.EmployeeAbbreviation,
		Description:          &c.Description,
	}
}

func ComputerFromDomain(comp domain.Computer) computer {
	compStorage := computer{Name: comp.Name}
	if comp.MAC != nil {
		compStorage.MAC = *comp.MAC
	}
	if comp.IP != nil {
		compStorage.IP = *comp.IP
	}
	if comp.EmployeeAbbreviation != nil {
		compStorage.EmployeeAbbreviation = *comp.EmployeeAbbreviation
	}
	if comp.Description != nil {
		compStorage.Description = *comp.Description
	}

	return compStorage
}
