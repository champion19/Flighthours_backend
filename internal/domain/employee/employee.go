package employee

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Employee struct {
	ID                   string
	Name                 string
	Airline              string
	Email                string
	Password             string
	Emailconfirmed       bool
	IdentificationNumber string
	Bp                   string
  StartDate           time.Time
	EndDate             time.Time
	Active              bool
}

func (e *Employee) setID() {
	e.ID = uuid.New().String()
}

func (e *Employee) hashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	e.Password = string(hash)
	return nil
}

type Airline struct {
	ID   string
	Name string
	Code string
	Status string
}
