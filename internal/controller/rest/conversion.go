package rest

import (
	"strings"

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

func ComputerRestModelToDomain(restModel models.Computer, propertiesToUpdate []string) domain.Computer {
	computerDomain := domain.Computer{}
	if restModel.Name != nil {
		computerDomain.Name = *restModel.Name
	}

	if len(propertiesToUpdate) == 0 {
		computerDomain.MAC = toPtr(restModel.Mac)
		computerDomain.IP = toPtr(restModel.IP)
		computerDomain.EmployeeAbbreviation = toPtr(restModel.EmployeeAbbreviation)
		computerDomain.Description = toPtr(restModel.Description)
	} else {
		for _, propertyCaseSensitive := range propertiesToUpdate {
			switch property := strings.ToLower(propertyCaseSensitive); property {
			case "mac":
				computerDomain.MAC = toPtr(restModel.Mac)
			case "ip":
				computerDomain.IP = toPtr(restModel.IP)
			case "employeeabbreviation":
				computerDomain.EmployeeAbbreviation = toPtr(restModel.EmployeeAbbreviation)
			case "description":
				computerDomain.Description = toPtr(restModel.Description)
			}
		}
	}

	return computerDomain
}

func NewComputerRestModelToDomain(restModel models.NewComputer) domain.Computer {
	return ComputerRestModelToDomain(restModel.Computer, nil)
}
