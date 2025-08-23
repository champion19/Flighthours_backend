// api/dependencies.go
package api

import (
	"fmt"
	"log"

	"github.com/champion19/Flighthours_backend/config"
	domain "github.com/champion19/Flighthours_backend/internal/domain/employee"
	repo "github.com/champion19/Flighthours_backend/internal/platform/employee"
	"github.com/champion19/Flighthours_backend/internal/platform/notification"
	"github.com/champion19/Flighthours_backend/internal/platform/jwt"
)

type Dependencies struct {
	employeeService domain.Service
	config          *config.Config
}

func initDependencies() *Dependencies {
	cfg := config.Load()

	db, err := repo.GetDB(cfg.Database)
	if err != nil {
		log.Printf("Database configuration: %+v", cfg.Database)
		log.Printf("Database connection error: %v", err)
		panic(fmt.Sprintf("error connecting to database: %v", err))
	}
	employeeRepo := repo.NewRepository(db)

	tokenGen := jwt.New(cfg.JWT.SecretKey)

	resendNotifier := notification.NewResendNotifier(cfg.Resend)

	employeeService := domain.NewService(employeeRepo, resendNotifier, tokenGen, cfg)

	return &Dependencies{
		employeeService: employeeService,
		config:          &cfg,
	}
}
