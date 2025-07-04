package api
import(
 "github.com/champion19/Flighthours_backend/config"
domain  "github.com/champion19/Flighthours_backend/internal/domain/employee"
repo  "github.com/champion19/Flighthours_backend/internal/platform/employee"
)

// ...

type Dependencies struct {
	employeeRepo domain.Repository
  config       *config.Config
}

func initDependencies() *Dependencies {
	cfg := config.Load()
	db, err := repo.GetDB(cfg.Database)
	if err != nil {
		panic("error get db")
	}
	employeeRepo := repo.NewRepository(db)

	return &Dependencies{
		employeeRepo: employeeRepo,
		config:       &cfg,
	}
}
