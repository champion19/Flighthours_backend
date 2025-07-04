package employee

import (
	"crypto/sha256"
	"encoding/base64"

	"github.com/google/uuid"
)

type Employee struct{
ID string
Name string
Email string
Password string
Emailconfirmed bool
IdentificationNumber int
Bp int
Fecha_inicio string
Fecha_fin string
Vigente bool

}

func (e *Employee) setID() {
	e.ID = uuid.New().String()
}

func (e *Employee) hashPassword() {
	hasher := sha256.New()
	hasher.Write([]byte(e.Password))
	e.Password = base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
