package app

import (
	"os"
	"time"

	"github.com/qrave1/lamoda_test/config"
	"github.com/qrave1/lamoda_test/internal/application"
	"github.com/qrave1/lamoda_test/internal/infrastructure/persistence"
	"github.com/qrave1/lamoda_test/internal/infrastructure/persistence/postgres"
	"github.com/qrave1/lamoda_test/internal/interface/http"
	"github.com/qrave1/lamoda_test/pkg/logger"
)

type App struct {
	srv *http.Server
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run(pathToConfig, pathToSpec string) {
	log := logger.New()

	cfg, err := config.ReadConfig(pathToConfig)
	if err != nil {
		log.Error("error reading config", "error", err)
		os.Exit(-1)
	}

	db, err := postgres.NewConnect(cfg)
	if err != nil {
		log.Error("error connecting to database", "error", err)
		os.Exit(-1)
	}

	// repositories
	productRepo := persistence.NewProductPostgresRepository(db, log)
	warehouseRepo := persistence.NewWarehousePostgresRepository(db, log)
	productWarehouseRepo := persistence.NewProductWarehousePostgresRepository(db, log)

	// services
	reservationService := application.NewReservationServiceImpl(db, warehouseRepo, productRepo, productWarehouseRepo, log)

	// api
	api := http.NewAPI(reservationService, log)

	// server
	srv := http.NewServer(cfg.Server.Port)

	app.srv = srv

	err = srv.Run(api, pathToSpec)
	if err != nil {
		log.Error("error starting server", "error", err)
		os.Exit(-1)
	}
}

func (app *App) Shutdown(timeout time.Duration) {
	app.srv.Shutdown(timeout)
}
