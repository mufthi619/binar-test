package database

import (
	"binar/pkg/config"
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB

type Databases struct {
	WriteDB *gorm.DB
	ReadDB  *gorm.DB
	TxDb    *gorm.DB
}

func NewConnection(cfg *config.Config) (*gorm.DB, error) {
	//Check Existence DB
	if db != nil {
		return db, nil
	}

	//Logger
	loggerConfig := logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger.Silent,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	}
	if cfg.AppConfig.DebugMode {
		loggerConfig.LogLevel = logger.Info
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		loggerConfig,
	)

	//GORM Config
	dbConfig := &gorm.Config{
		Logger: newLogger,
	}

	//GORM Open Conn
	var err error
	switch cfg.DatabaseConfig.UsageSQL {
	case config.MySQL:
		db, err = gorm.Open(mysql.Open(mysqlDSN(cfg.DatabaseConfig)), dbConfig)
		break
	case config.PgSQL:
		db, err = gorm.Open(postgres.Open(pgsqlDSN(cfg.DatabaseConfig)), dbConfig)
		break
	default:
		return nil, errors.New("invalid usage db")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	//Connection Pooling
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm.DB: %v", err)
	}
	sqlDB.SetMaxIdleConns(cfg.DatabaseConfig.MaxIdleConn)
	sqlDB.SetMaxOpenConns(cfg.DatabaseConfig.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.DatabaseConfig.MaxLifeTimeConn) * time.Second)

	//Ping to test the connection
	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		return nil, fmt.Errorf("database connection ping failed: %w", err)
	}

	return db, nil
}

func mysqlDSN(dbConfig config.DatabaseConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Database,
	)
}

func pgsqlDSN(dbConfig config.DatabaseConfig) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Database,
	)
}
