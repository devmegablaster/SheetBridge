package database

import (
	"fmt"
	"log/slog"

	"github.com/devmegablaster/SheetBridge/internal/config"
	"github.com/devmegablaster/SheetBridge/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseSvc struct {
	DB *gorm.DB
}

func New(c *config.DatabaseConfig) *DatabaseSvc {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		c.Host, c.User, c.Password, c.Name, c.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// INFO: Disable gorm logging
	db.Logger = db.Logger.LogMode(0)

	uuidExtension(db)
	autoMigrate(db)

	if err != nil {
		panic(err)
	}

	slog.Info("✅Database connected")

	return &DatabaseSvc{
		DB: db,
	}
}

func autoMigrate(conn *gorm.DB) {
	conn.AutoMigrate(&models.User{}, &models.Connection{}, &models.DatabaseConfig{}, &models.Synk{}, &models.Schema{})
}

func uuidExtension(conn *gorm.DB) {
	conn.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
}
