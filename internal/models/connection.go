package models

import (
	"time"

	"github.com/google/uuid"
)

type Connection struct {
	Id               uuid.UUID      `json:"id" gorm:"primary_key;default:uuid_generate_v4();type:uuid"`
	DatabaseConfigId uuid.UUID      `json:"databaseConfigId" gorm:"type:uuid;not null"`
	DatabaseConfig   DatabaseConfig `json:"databaseConfig" gorm:"foreignKey:DatabaseConfigId"`
	Tables           []string       `json:"tables" gorm:"-"`
	UserId           uuid.UUID      `json:"userId" gorm:"type:uuid;not null"`
	CreatedAt        time.Time      `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `json:"updatedAt" gorm:"autoUpdateTime"`
}

type DatabaseConfig struct {
	Id       uuid.UUID `json:"id" gorm:"primary_key;default:uuid_generate_v4();type:uuid"`
	Host     string    `json:"host" validate:"required"`
	Username string    `json:"username" validate:"required"`
	Password string    `json:"password" validate:"required"`
	Database string    `json:"database" validate:"required"`
	Port     string    `json:"port" validate:"required"`
}

type ConnectionRequest struct {
	DatabaseConfig DatabaseConfig `json:"databaseConfig" validate:"required"`
}

func (c *ConnectionRequest) ToConnection(userId uuid.UUID) Connection {
	return Connection{
		DatabaseConfig: c.DatabaseConfig,
		UserId:         userId,
	}
}

type ConnectionResponse struct {
	Id             uuid.UUID              `json:"id" gorm:"primary_key;default:uuid_generate_v4();type:uuid"`
	DatabaseConfig DatabaseConfigResponse `json:"databaseConfig"`
	Tables         []string               `json:"tables"`
}

type DatabaseConfigResponse struct {
	Id   string `json:"id"`
	Host string `json:"host"`
}

func (c *Connection) ToResponse() ConnectionResponse {
	return ConnectionResponse{
		Id:     c.Id,
		Tables: c.Tables,
		DatabaseConfig: DatabaseConfigResponse{
			Id:   c.DatabaseConfig.Id.String(),
			Host: c.DatabaseConfig.Host,
		},
	}
}
