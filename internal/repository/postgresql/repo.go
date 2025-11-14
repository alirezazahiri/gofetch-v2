package postgresql

import (
	"fmt"
	"log"
	"time"

	"github.com/alirezazahiri/gofetch-v2/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host     string `koanf:"host"`
	Port     uint   `koanf:"port"`
	User     string `koanf:"username"`
	Password string `koanf:"password"`
	DBName   string `koanf:"dbname"`
}

type Repository struct {
	db *gorm.DB
}

func New(cfg *Config) (*Repository, error) {
	// First, connect to the default 'postgres' database to create our target database if it doesn't exist
	defaultDSN := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=postgres sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password,
	)

	defaultDB, err := gorm.Open(postgres.Open(defaultDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to default database: %w", err)
	}

	// Create the target database if it doesn't exist
	createDBSQL := fmt.Sprintf("CREATE DATABASE \"%s\"", cfg.DBName)
	result := defaultDB.Exec(createDBSQL)
	if result.Error != nil {
		// Check if error is because database already exists (which is fine)
		log.Printf("Note: %v (this is fine if database already exists)", result.Error)
	} else {
		log.Printf("Database '%s' created successfully", cfg.DBName)
	}

	// Close the default database connection
	sqlDefaultDB, err := defaultDB.DB()
	if err == nil {
		sqlDefaultDB.Close()
	}

	// Now connect to our target database
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database")

	// Auto-migrate the schema
	if err := db.AutoMigrate(&entity.Job{}, &entity.JobResult{}); err != nil {
		return nil, fmt.Errorf("failed to auto-migrate database schema: %w", err)
	}

	log.Println("Database schema migrated successfully")

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) DB() *gorm.DB {
	return r.db
}

func (r *Repository) Close() error {
	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
