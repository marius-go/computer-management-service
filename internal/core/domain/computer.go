package domain

import "fmt"

const LenEmployeeAbbreviation int = 3

type Computer struct {
	Name                 string
	MAC                  *string
	IP                   *string
	EmployeeAbbreviation *string
	Description          *string
}

func (c *Computer) Validate() error {
	if c.Name == "" {
		return ErrValidation("property 'Name' is required")
	}
	if c.MAC == nil || *c.MAC == "" {
		return ErrValidation("property 'MAC' is required")
	}
	if c.IP == nil || *c.IP == "" {
		return ErrValidation("property 'IP' is required")
	}
	if c.EmployeeAbbreviation != nil && len(*c.EmployeeAbbreviation) > 0 && len(*c.EmployeeAbbreviation) != LenEmployeeAbbreviation {
		return ErrValidation(fmt.Sprint("if set, 'EmployeeAbbreviation' must be exactly ", LenEmployeeAbbreviation, " characters long"))
	}

	return nil
}
