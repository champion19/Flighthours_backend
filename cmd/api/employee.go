package api

import (
	"fmt"
	"html/template"

	domain "github.com/champion19/Flighthours_backend/internal/domain/employee"
	"github.com/go-playground/validator/v10"
)

type EmployeeRequest struct {
	Name                 string `json:"Name"`
	Airline              string `json:"Airline"`
	Email                string `json:"Email"`
	Password             string `json:"Password"`
	Emailconfirmed       bool   `json:"Emailconfirmed"`
	IdentificationNumber string `json:"IdentificationNumber"`
	Bp                   string `json:"Bp"`
	StartDate            string `json:"Start_date"`
	EndDate              string `json:"End_date"`
	Active               bool   `json:"Active"`
}

type EmployeeLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (e EmployeeLogin) Validate() any {
	panic("unimplemented")
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type EmployeeResponse struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Airline              string `json:"airline"`
	Email                string `json:"email"`
	Emailconfirmed       bool   `json:"emailconfirmed"`
	IdentificationNumber string `json:"identification_number"`
	Bp                   string `json:"bp"`
	StartDate            string `json:"start_date"`
	EndDate              string `json:"end_date"`
	Active               bool   `json:"active"`

}

type LoginResponse struct {
	ID                   string `json:"id"`
	Name                 string `json:"name"`
	Airline              string `json:"airline"`
	Email                string `json:"email"`
	Emailconfirmed       bool   `json:"emailconfirmed"`
	IdentificationNumber string `json:"identification_number"`
	Bp                   string `json:"bp"`
	StartDate            string `json:"start_date"`
	EndDate              string `json:"end_date"`
	Active               bool   `json:"active"`
	Token                string `json:"token"`
}
type ResponseEmail struct {
	Title   string
	Content template.HTML
}

type AirlineResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type EmployeeEmailConfirmedResponse struct {
	Emailconfirmed bool `json:"emailconfirmed"`
}

func (e EmployeeLogin) ToDomain() domain.Employee {
	return domain.Employee{
		Email:    e.Email,
		Password: e.Password,
	}
}

func (e LoginRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(e)
}

func (e EmployeeRequest) ToDomain() domain.Employee {
	return domain.Employee{
		Name:                 e.Name,
		Airline:              e.Airline,
		Email:                e.Email,
		Password:             e.Password,
		Emailconfirmed:       e.Emailconfirmed,
		IdentificationNumber: e.IdentificationNumber,
		Bp:                   e.Bp,
		StartDate:            e.StartDate,
		EndDate:              e.EndDate,
		Active:               e.Active,
	}
}

func (e LoginRequest) ToDomain() domain.Employee {
	return domain.Employee{
		Email:    e.Email,
		Password: e.Password,
	}
}

func (e EmployeeRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(e)
	if err != nil {
		validateErrors := err.(validator.ValidationErrors)
		message := ""

		for _, validateErr := range validateErrors {
			message += fmt.Sprintf("%s: %s,", validateErr.Field(), validateErr.Error())
		}

		return fmt.Errorf("%w: %s", ErrValidationUser, message)
	}
	return nil
}
