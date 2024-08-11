package main

import (
	"binar/internal/infra/database"
	"binar/internal/seeder"
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func newMigrateCommand(logger *zap.Logger, db *gorm.DB) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Run database migrations",
		RunE: func(cmd *cobra.Command, args []string) error {
			migrationRunner := database.NewMigrationRunner("./migrations", db, logger)
			if err := migrationRunner.Run(); err != nil {
				return fmt.Errorf("error running migrations: %v", err)
			}
			logger.Info("Migrations completed successfully")
			return nil
		},
	}

	cmd.AddCommand(
		&cobra.Command{
			Use:   "up",
			Short: "Run all up migrations",
			RunE: func(cmd *cobra.Command, args []string) error {
				migrationRunner := database.NewMigrationRunner("./migrations", db, logger)
				if err := migrationRunner.Run(); err != nil {
					return fmt.Errorf("error running up migrations: %v", err)
				}
				logger.Info("Up migrations completed successfully")
				return nil
			},
		},
		&cobra.Command{
			Use:   "down",
			Short: "Run all down migrations",
			RunE: func(cmd *cobra.Command, args []string) error {
				migrationRunner := database.NewMigrationRunner("./migrations", db, logger)
				if err := migrationRunner.Rollback(); err != nil {
					return fmt.Errorf("error running down migrations: %v", err)
				}
				logger.Info("Down migrations completed successfully")
				return nil
			},
		},
	)

	return cmd
}

func newSeedCommand(logger *zap.Logger, db *gorm.DB) *cobra.Command {
	return &cobra.Command{
		Use:   "seed",
		Short: "Run database seeder",
		RunE: func(cmd *cobra.Command, args []string) error {
			seederRun := seeder.NewUserSeeder(db, logger)
			if err := seederRun.Seed(); err != nil {
				return fmt.Errorf("error seeding database: %v", err)
			}
			logger.Info("Database seeded successfully")
			return nil
		},
	}
}
