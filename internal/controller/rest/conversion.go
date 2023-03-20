package rest

import (
	"github.com/marius-go/computer-management-service/gen/controller/rest/models"
	"github.com/marius-go/computer-management-service/internal/core/domain"
)

func ComputerRestModelFromDomain(computerDomain domain.Computer) models.Computer {
	restModel := models.Computer{Name: &computerDomain.Name}

	if computerDomain.MAC != nil {
		restModel.Mac = *computerDomain.MAC
	}

	if computerDomain.IP != nil {
		restModel.IP = *computerDomain.IP
	}

	if computerDomain.EmployeeAbbreviation != nil {
		restModel.EmployeeAbbreviation = *computerDomain.EmployeeAbbreviation
	}

	if computerDomain.Description != nil {
		restModel.Description = *computerDomain.Description
	}

	return restModel
}

func toPtr[T any](value T) *T {
	return &value
}

func ComputerRestModelToDomain(restModel models.Computer) domain.Computer {
	computerDomain := domain.Computer{
		MAC:                  toPtr(restModel.Mac),
		IP:                   toPtr(restModel.IP),
		EmployeeAbbreviation: toPtr(restModel.EmployeeAbbreviation),
		Description:          toPtr(restModel.Description),
	}

	if restModel.Name != nil {
		computerDomain.Name = *restModel.Name
	}

	return computerDomain
}

func NewComputerRestModelToDomain(restModel models.NewComputer) domain.Computer {
	return ComputerRestModelToDomain(restModel.Computer)
}
