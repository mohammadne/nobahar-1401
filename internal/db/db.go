package db

import (
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB interface {
	Migrate(ctx context.Context) error
}

type db struct {
	instance *gorm.DB
}

func New(cfg *Config) (*db, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database, cfg.SSLMode, cfg.Timezone,
	)

	gorm, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &db{gorm}, nil
}

func (db *db) Migrate(ctx context.Context) error {
	return db.instance.AutoMigrate()
}
