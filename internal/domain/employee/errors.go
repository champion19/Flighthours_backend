package employee

import "errors"

var (
	ErrEmployeeCannotSave      = errors.New("error employee can not save")
	ErrGetEmployee             = errors.New("error get employee")
	ErrDuplicateEmployee       = errors.New("employee already exists")
	ErrSavingEmployee          = errors.New("error saving employee")
	ErrEmployeeCannotGet       = errors.New("error can not get user")
	ErrEmployeeCannotFound     = errors.New("error can no found user")
	ErrGettingEmployeeByEmail  = errors.New("error getting user by the email")
	ErrNotFoundEmployeeByEmail = errors.New("error not found user by email")
	ErrEmployeeCannotLogin     = errors.New("error user can not login")
	ErrValidationEmployee      = errors.New("error validation user")
	ErrInvalidJson             = errors.New("error invalid json")
	ErrInvalidToken            = errors.New("error invalid token")
	ErrEmailAlreadyVerified    = errors.New("email already verified")
	ErrEmailNotVerified        = errors.New("email not verified")
	ErrTokenExpired           = errors.New("token expired")
	ErrTokenNotFound          = errors.New("token not found")
	ErrTokenInvalid     = errors.New("token invalid")
	ErrTokenMalformed   = errors.New("token malformed")
	ErrTokenUnexpected  = errors.New("unexpected signing method")
	ErrAirlineCannotFound = errors.New("airline not found")
	ErrAirlineCannotGet = errors.New("airline cannot get")
  ErrGetAirline = errors.New("error get airline")
)
