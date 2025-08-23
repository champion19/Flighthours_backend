package employee

import (
	"fmt"
	"log"
	"time"

	"github.com/champion19/Flighthours_backend/config"

	"github.com/champion19/Flighthours_backend/internal/domain/token"
	"github.com/champion19/Flighthours_backend/internal/platform/notification"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	Save(employee Employee) error
	GetEmployeeByID(id string) (*Employee, error)
	GetEmployeeByEmail(email string) (*Employee, error)
	UpdateEmailConfirmed(employeeID string, confirmed bool) error
	GetAllAirlines() ([]Airline, error)
}

type Service interface {
	Save(employee Employee) (Employee, error)
	GetEmployeeByID(id string) (*Employee, error)
	GetEmployeeByEmail(email string) (*Employee, error)
	Login(employee Employee) (*Employee, string, error)
	VerifyEmail(token string) error
	GetAllAirlines() ([]Airline, error)
}

type service struct {
	repository         Repository
	notificationSender notification.ResendNotifier
	tokenGenerator     token.Generator
	cfg                config.Config
}

func NewService(repo Repository, notificationSender notification.ResendNotifier, tokenGen token.Generator, cfg config.Config) Service {
	return &service{
		repository:         repo,
		notificationSender: notificationSender,
		tokenGenerator:     tokenGen,
		cfg:                cfg,
	}
}

func (s service) GetEmployeeByID(id string) (*Employee, error) {
	return s.repository.GetEmployeeByID(id)
}

func (s service) GetEmployeeByEmail(email string) (*Employee, error) {
	return s.repository.GetEmployeeByEmail(email)
}

func (s service) Save(employee Employee) (Employee, error) {
	existingEmployee, err := s.GetEmployeeByEmail(employee.Email)
	if err == nil && existingEmployee != nil {
		return Employee{}, ErrDuplicateEmployee
	}

	employee.setID()
	if err := employee.hashPassword(); err != nil {
		return Employee{}, err
	}
	err = s.repository.Save(employee)
	if err != nil {
		return Employee{}, err
	}

	go func() {
		err := s.sendVerificationEmail(employee.ID, employee.Email, s.cfg.VerificationToken.ExpirationTime)
		if err != nil {
			log.Printf("could not send verification email: %v", err)
		}
	}()

	return employee, nil
}

func (s service) sendVerificationEmail(employeeID, email string, duration time.Duration) error {

	token, err := s.tokenGenerator.GenerateJWT(employeeID, duration)
	if err != nil {
		return fmt.Errorf("failed to generate verification token: %w", err)
	}

	verificationLink := fmt.Sprintf("%s/v1/employees/verify-email?token=%s", s.cfg.API.BaseURL, token)

	return s.notificationSender.SendVerificationEmail(email, verificationLink)
}

func (s service) VerifyEmail(tokenStr string) error {

	employeeID, err := s.tokenGenerator.ValidateJWT(tokenStr)
	if err != nil {
		return err
	}

	employee, err := s.repository.GetEmployeeByID(employeeID)
	if err != nil {
		return ErrEmployeeCannotGet
	}

	if employee.Emailconfirmed {
		return ErrEmailAlreadyVerified
	}

	return s.repository.UpdateEmailConfirmed(employeeID, true)
}

func (e Employee) comparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(e.Password), []byte(password))
	return err
}

func (s service) Login(employee Employee) (*Employee, string, error) {
	employeeFound, err := s.GetEmployeeByEmail(employee.Email)
	if err != nil {
		return nil, "", err
	}

	if !employeeFound.Emailconfirmed {
		return nil, "", ErrEmailNotVerified
	}

	if err := employeeFound.comparePassword(employee.Password); err != nil {
		return nil, "", err
	}
	cfg := config.Load()
	token, err := s.tokenGenerator.GenerateJWT(
		employeeFound.ID,
		cfg.JWT.ExpirationTime)
	if err != nil {
		return nil, "", err
	}

	return employeeFound, token, nil
}


func (s *service) GetAllAirlines() ([]Airline, error) {
	return s.repository.GetAllAirlines()
}





