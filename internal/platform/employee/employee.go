package employee

import (
	domain "github.com/champion19/Flighthours_backend/internal/domain/employee"
)

type Employee struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	Emailconfirmed       bool   `json:"emailconfirmed"`
	IdentificationNumber int    `json:"identificationNumber"`
	Bp                   int    `json:"bp"`
	Fecha_inicio         string `json:"fecha_inicio"`
	Fecha_fin            string `json:"fecha_fin"`
	Vigente              bool   `json:"vigente"`
}

func (e Employee) ToDomain() domain.Employee {
	return domain.Employee{
		ID:                   e.ID,
		Name:                 e.Name,
		Email:                e.Email,
		Password:             e.Password,
		Emailconfirmed:       e.Emailconfirmed,
		IdentificationNumber: e.IdentificationNumber,
		Bp:                   e.Bp,
		Fecha_inicio:         e.Fecha_inicio,
		Fecha_fin:            e.Fecha_fin,
		Vigente:              e.Vigente,
	}
}
