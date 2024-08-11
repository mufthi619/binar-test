package main

import (
	"binar/pkg/config"
	"binar/pkg/di"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cfg, err := config.LoadConfig("config/app.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	infra, err := di.InitializeApp(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	rootCmd := &cobra.Command{Use: "boilerplate"}
	rootCmd.AddCommand(
		newMigrateCommand(infra.Logger, infra.Databases.WriteDB),
		newSeedCommand(infra.Logger, infra.Databases.WriteDB),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
