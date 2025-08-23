package token

import (
	"time"
)
type Generator interface {
	GenerateJWT(employeeID string,duration time.Duration) (string, error)
	ValidateJWT(token string) (employeeID string, err error)
}
