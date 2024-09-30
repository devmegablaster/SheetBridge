package repository

import (
	"github.com/devmegablaster/SheetBridge/internal/database"
	"github.com/devmegablaster/SheetBridge/internal/models"
	"github.com/google/uuid"
)

type ConnectionRepository struct {
	dbSvc *database.DatabaseSvc
}

func NewConnectionRepository(dbSvc *database.DatabaseSvc) *ConnectionRepository {
	return &ConnectionRepository{
		dbSvc: dbSvc,
	}
}

func (r *ConnectionRepository) CreateConnection(c *models.Connection) error {
	return r.dbSvc.DB.Create(c).Error
}

func (r *ConnectionRepository) GetConnections() ([]models.Connection, error) {
	var connections []models.Connection
	err := r.dbSvc.DB.Preload("DatabaseConfig").Find(&connections).Error
	return connections, err
}

func (r *ConnectionRepository) GetConnectionById(id uuid.UUID) (*models.Connection, error) {
	var connection models.Connection
	err := r.dbSvc.DB.Preload("DatabaseConfig").Where("id = ?", id).First(&connection).Error
	return &connection, err
}

func (r *ConnectionRepository) GetConnectionsByUserId(id uuid.UUID) ([]models.Connection, error) {
	var connections []models.Connection
	err := r.dbSvc.DB.Preload("DatabaseConfig").Where("user_id = ?", id).Find(&connections).Error
	return connections, err
}
