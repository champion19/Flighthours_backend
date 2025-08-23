package employee

import (


	domain "github.com/champion19/Flighthours_backend/internal/domain/employee"
)

type Employee struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Airline              string `json:"airline"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	Emailconfirmed       bool   `json:"emailconfirmed"`
	IdentificationNumber string    `json:"identification_number"`
	Bp                   string    `json:"bp"`
	StartDate           string `json:"start_date"`
	EndDate            string `json:"end_date"`
	Active              bool   `json:"active"`
}

func (e Employee) ToDomain() domain.Employee {
	return domain.Employee{
		ID:                   e.ID,
		Name:                 e.Name,
		Email:                e.Email,
		Airline:              e.Airline,
		Password:             e.Password,
		Emailconfirmed:       e.Emailconfirmed,
		IdentificationNumber: e.IdentificationNumber,
		Bp:                   e.Bp,
	  StartDate:         e.StartDate,
		EndDate:            e.EndDate,
		Active:              e.Active,
	}
}

type Airline struct {
	ID   string
	Name string
	Code string
	Status string
}

func (a Airline) ToDomain() domain.Airline {
	return domain.Airline{
		ID:                   a.ID,
		Name:                 a.Name,
		Code:                 a.Code,
		Status:               a.Status,

	}
}
