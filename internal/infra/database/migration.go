package database

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MigrationRunner struct {
	logger         *zap.Logger
	migrationsPath string
	db             *gorm.DB
}

func NewMigrationRunner(migrationsPath string, db *gorm.DB, logger *zap.Logger) *MigrationRunner {
	return &MigrationRunner{
		migrationsPath: migrationsPath,
		db:             db,
		logger:         logger,
	}
}

func (r *MigrationRunner) Run() error {
	m, err := r.createMigrateInstance()
	if err != nil {
		return fmt.Errorf("error creating migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error running migrations: %v", err)
	}

	r.logger.Info("Migrations completed successfully")
	return nil
}

func (r *MigrationRunner) Rollback() error {
	m, err := r.createMigrateInstance()
	if err != nil {
		return fmt.Errorf("error creating migrate instance: %v", err)
	}

	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("error rolling back migrations: %v", err)
	}

	r.logger.Info("Rollback completed successfully")
	return nil
}

func (r *MigrationRunner) createMigrateInstance() (*migrate.Migrate, error) {
	sqlDB, err := r.db.DB()
	if err != nil {
		return nil, fmt.Errorf("error getting underlying sql.DB: %v", err)
	}

	var driver database.Driver
	switch r.db.Dialector.Name() {
	case "postgres":
		driver, err = postgres.WithInstance(sqlDB, &postgres.Config{})
	case "mysql":
		driver, err = mysql.WithInstance(sqlDB, &mysql.Config{})
	default:
		return nil, fmt.Errorf("unsupported database type: %s", r.db.Dialector.Name())
	}

	if err != nil {
		return nil, fmt.Errorf("error creating database driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", r.migrationsPath),
		r.db.Dialector.Name(),
		driver,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating migrate instance: %v", err)
	}

	return m, nil
}
