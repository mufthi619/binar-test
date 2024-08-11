package di

import (
	"binar/internal/infra"
	"binar/internal/infra/database"
	"binar/internal/infra/logger"
	"binar/internal/infra/queue"
	"binar/pkg/config"
	"github.com/google/wire"
)

var InfraSet = wire.NewSet(
	ProvideDatabases,
	queue.NewRabbitMQConn,
	logger.NewZap,
	wire.Struct(new(infra.Infra), "*"),
)

func ProvideDatabases(cfg *config.Config) (*database.Databases, error) {
	writeDB, err := database.NewConnection(cfg)
	if err != nil {
		return nil, err
	}
	readDB, err := database.NewConnection(cfg)
	if err != nil {
		return nil, err
	}
	return &database.Databases{
		WriteDB: writeDB,
		ReadDB:  readDB,
	}, nil
}
