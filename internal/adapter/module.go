package adapter

import (
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/http/external/repository"
	"github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase1/internal/adapter/storage/postgres"
	"go.uber.org/fx"
)

var Module = fx.Options(
	http.Module,
	postgres.Module,
	repository.Module,
)
