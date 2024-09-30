package services

import (
	"github.com/devmegablaster/SheetBridge/internal/config"
	"github.com/devmegablaster/SheetBridge/internal/database"
	"github.com/devmegablaster/SheetBridge/internal/models"
	"github.com/devmegablaster/SheetBridge/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ConnectionService struct {
	cr            *repository.ConnectionRepository
	dbSvc         *database.DatabaseSvc
	encryptionSvc *EncryptionService
	validator     *validator.Validate
}

func NewConnectionService(dbSvc *database.DatabaseSvc, cfg config.CryptoConfig) *ConnectionService {
	return &ConnectionService{
		cr:            repository.NewConnectionRepository(dbSvc),
		dbSvc:         dbSvc,
		encryptionSvc: NewEncryptionService(cfg),
		validator:     validator.New(),
	}
}

func (s *ConnectionService) CreateConnection(c *models.Connection) error {
	if err := s.validator.Struct(c); err != nil {
		return err
	}

	// TODO: Health check and get tables

	// INFO: DatabaseConfig encryption
	var err error
	c.DatabaseConfig.Password, err = s.encryptionSvc.Encrypt(c.DatabaseConfig.Password)
	c.DatabaseConfig.Username, err = s.encryptionSvc.Encrypt(c.DatabaseConfig.Username)
	c.DatabaseConfig.Host, err = s.encryptionSvc.Encrypt(c.DatabaseConfig.Host)
	c.DatabaseConfig.Port, err = s.encryptionSvc.Encrypt(c.DatabaseConfig.Port)
	c.DatabaseConfig.Database, err = s.encryptionSvc.Encrypt(c.DatabaseConfig.Database)
	if err != nil {
		return err
	}

	return s.cr.CreateConnection(c)
}

func (s *ConnectionService) GetConnectionById(id uuid.UUID) *models.Connection {
	conn, err := s.cr.GetConnectionById(id)
	if err != nil {
		return nil
	}

	return conn
}

func (s *ConnectionService) GetConnectionsForUser(user *models.User) ([]models.ConnectionResponse, error) {
	connections, err := s.cr.GetConnectionsByUserId(user.Id)
	if err != nil {
		return nil, err
	}

	connectionResponse := []models.ConnectionResponse{}
	for _, c := range connections {
		c.DatabaseConfig.Host, _ = s.encryptionSvc.Decrypt(c.DatabaseConfig.Host)
		connectionResponse = append(connectionResponse, c.ToResponse())
	}

	return connectionResponse, nil
}

func (s *ConnectionService) GetConnections() ([]models.ConnectionResponse, error) {
	connections, err := s.cr.GetConnections()
	if err != nil {
		return nil, err
	}

	connectionResponse := []models.ConnectionResponse{}
	for _, c := range connections {
		c.DatabaseConfig.Host, _ = s.encryptionSvc.Decrypt(c.DatabaseConfig.Host)
		connectionResponse = append(connectionResponse, c.ToResponse())
	}

	return connectionResponse, nil
}

func (s *ConnectionService) DecryptConnection(c *models.Connection) *models.Connection {
	c.DatabaseConfig.Host, _ = s.encryptionSvc.Decrypt(c.DatabaseConfig.Host)
	c.DatabaseConfig.Port, _ = s.encryptionSvc.Decrypt(c.DatabaseConfig.Port)
	c.DatabaseConfig.Username, _ = s.encryptionSvc.Decrypt(c.DatabaseConfig.Username)
	c.DatabaseConfig.Password, _ = s.encryptionSvc.Decrypt(c.DatabaseConfig.Password)
	c.DatabaseConfig.Database, _ = s.encryptionSvc.Decrypt(c.DatabaseConfig.Database)
	return c
}
