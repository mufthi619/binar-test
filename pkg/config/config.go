package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type UsageSQL string

const (
	PgSQL UsageSQL = "pgsql"
	MySQL UsageSQL = "mysql"
)

type (
	Config struct {
		AppConfig      AppConfig      `yaml:"app_config"`
		DatabaseConfig DatabaseConfig `yaml:"database_config"`
		RabbitConfig   RabbitMQ       `yaml:"rabbit_config"`
	}

	AppConfig struct {
		AppName   string `yaml:"app_name"`
		Port      int    `yaml:"port"`
		DebugMode bool   `yaml:"debug_mode"`
		URL       string `yaml:"url"`
		PublicDir string `yaml:"public_dir"`
	}

	DatabaseConfig struct {
		Host            string   `yaml:"host"`
		Port            int      `yaml:"port"`
		User            string   `yaml:"user"`
		Password        string   `yaml:"password"`
		Database        string   `yaml:"database"`
		UsageSQL        UsageSQL `yaml:"usage_sql"`
		MaxIdleConn     int      `yaml:"max_idle_conn"`
		MaxOpenConn     int      `yaml:"max_open_conn"`
		MaxLifeTimeConn int      `yaml:"max_life_time_conn"`
	}

	RabbitMQ struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
	}
)

func LoadConfig(path string) (*Config, error) {
	var cfg Config
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config file: %w", err)
	}

	return &cfg, nil
}
