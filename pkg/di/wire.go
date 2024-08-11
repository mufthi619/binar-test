//go:build wireinject
// +build wireinject

package di

import (
	"binar/internal/infra"
	"binar/pkg/config"
	"github.com/google/wire"
)

func InitializeApp(cfg *config.Config) (*infra.Infra, error) {
	wire.Build(
		ProvideAppConfig,
		InfraSet,
		RepositorySet,
		ServiceSet,
		HandlerSet,
	)
	return &infra.Infra{}, nil
}

func ProvideAppConfig(cfg *config.Config) *config.AppConfig {
	return &cfg.AppConfig
}
