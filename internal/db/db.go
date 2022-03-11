package db

import (
	"fmt"

	"github.com/mohammadne/nobahar-1401/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB interface {
	Migrate() error
	GetUser(email string) (*models.User, error)
	CreateUser(user *models.User) error
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

func (db *db) Migrate() error {
	return db.instance.AutoMigrate(models.User{}, models.Role{}, models.Group{}, models.Chat{}, models.Message{})
}

// ===========================================================================> USER

func (db *db) GetUser(email string) (*models.User, error) {
	user := models.User{}
	result := db.instance.First(&user, "email = ?", email)
	return &user, result.Error
}

func (db *db) CreateUser(user *models.User) error {
	result := db.instance.Create(user)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
