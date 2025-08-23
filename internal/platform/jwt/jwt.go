package jwt

import (
	"errors"
	"time"

	domain "github.com/champion19/Flighthours_backend/internal/domain/employee"
	"github.com/champion19/Flighthours_backend/internal/domain/token"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type jwtGenerator struct {
	secretKey []byte
}

func New(secret string) token.Generator {
	return &jwtGenerator{
		secretKey: []byte(secret),
	}
}

func (j *jwtGenerator) GenerateJWT(employeeID string, duration time.Duration) (string, error) {
	claims := &Claims{
		ID: employeeID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   employeeID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *jwtGenerator) ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenSignatureInvalid
		}
		return j.secretKey, nil
	})

	if err != nil {
				switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return "", domain.ErrTokenExpired
		case errors.Is(err, jwt.ErrTokenMalformed):
			return "", domain.ErrTokenMalformed
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return "", domain.ErrTokenUnexpected
		case errors.Is(err, jwt.ErrTokenUnverifiable):
			return "", domain.ErrTokenInvalid
		case errors.Is(err, jwt.ErrTokenNotValidYet):
			return "", domain.ErrTokenInvalid
		case errors.Is(err, jwt.ErrTokenUsedBeforeIssued):
			return "", domain.ErrTokenInvalid
		case errors.Is(err, jwt.ErrTokenInvalidAudience):
			return "", domain.ErrTokenInvalid
		case errors.Is(err, jwt.ErrTokenInvalidIssuer):
			return "", domain.ErrTokenInvalid
		case errors.Is(err, jwt.ErrTokenInvalidSubject):
			return "", domain.ErrTokenInvalid
		case errors.Is(err, jwt.ErrTokenInvalidId):
			return "", domain.ErrTokenInvalid
		case errors.Is(err, jwt.ErrTokenInvalidClaims):
			return "", domain.ErrTokenInvalid
		case errors.Is(err, jwt.ErrTokenRequiredClaimMissing):
			return "", domain.ErrTokenInvalid
		case errors.Is(err, jwt.ErrInvalidKey):
			return "", domain.ErrTokenInvalid
		case errors.Is(err, jwt.ErrInvalidKeyType):
			return "", domain.ErrTokenInvalid
		case errors.Is(err, jwt.ErrHashUnavailable):
			return "", domain.ErrTokenInvalid
		case errors.Is(err, jwt.ErrInvalidType):
			return "", domain.ErrTokenInvalid
		default:
			return "", domain.ErrTokenInvalid
		}
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.ID, nil
	}

	return "", domain.ErrTokenInvalid
}
