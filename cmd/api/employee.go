package api

import (
	"fmt"

	domain "github.com/champion19/Flighthours_backend/internal/domain/employee"
	"github.com/go-playground/validator/v10"
)



type EmployeeRequest struct {
	Name     string `json:"name" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type EmployeeLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type EmployeeLoginResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Emailconfirmed bool   `json:"emailconfirmed"`
	IdentificationNumber int `json:"identificationNumber"`
	Bp int `json:"bp"`
	Fecha_inicio string `json:"fecha_inicio"`
	Fecha_fin string `json:"fecha_fin"`
	Vigente bool `json:"vigente"`
}


type EmployeeResponse struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
	Email string `json:"email"`
	Password string `json:"password"`
	Emailconfirmed bool   `json:"emailconfirmed"`
	IdentificationNumber int `json:"identificationNumber"`
	Bp int `json:"bp"`
	Fecha_inicio string `json:"fecha_inicio"`
	Fecha_fin string `json:"fecha_fin"`
	Vigente bool `json:"vigente"`
}


func (e EmployeeRequest) ToDomain() domain.Employee {
	return domain.Employee{
		Name:     e.Name,
		Email:    e.Email,
		Password: e.Password,
	}
}

func (e EmployeeLogin) ToDomain() domain.Employee {
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



func (e EmployeeLogin) Validate() error {
	validate := validator.New()
	err := validate.Struct(e)
	if err != nil {
		validateErrors := err.(validator.ValidationErrors)
		message := ""

		for _, validateErr := range validateErrors {
			message += fmt.Sprintf("%s: %s,", validateErr.Field(), validateErr.Error())
		}
		return fmt.Errorf(ErrValidationUser.Error(), message)
	}
	return nil
}
