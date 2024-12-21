package postgres

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/config"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/storage/postgres/repository"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"log/slog"
	"os"
)

func NewDatabaseConnection(cfg *config.Environment) *gorm.DB {
	// init database connection
	dbConnection, err := New(cfg.DatabaseURL)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}

	// migrate database
	if err = dbConnection.Migrate(); err != nil {
		slog.Error("error migrating database", "error", err)
		os.Exit(1)
	}
	return dbConnection.DB
}

var Module = fx.Options(
	fx.Provide(NewDatabaseConnection),
	repository.Module,
)
