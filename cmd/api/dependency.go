package api

import (
	 "github.com/champion19/Flighthours_backend/config"
	repo "github.com/champion19/Flighthours_backend/internal/platform/employee"
	domain"github.com/champion19/Flighthours_backend/internal/domain/employee"
)

type Dependencies struct {
	employee domain.Repository
}

func initDependencies() *Dependencies {
	cfg := config.Load()
	db, err := repo.GetDB(cfg.Database)
	if err != nil {
		panic("error get db")
	}
	employeerRepo := repo.NewRepository(db)

	return &Dependencies{
		employee: employeerRepo,
	}
}
