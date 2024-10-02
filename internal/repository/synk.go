package repository

import (
	"github.com/devmegablaster/SheetBridge/internal/database"
	"github.com/devmegablaster/SheetBridge/internal/models"
	"github.com/google/uuid"
)

type SynkRepository struct {
	dbSvc *database.DatabaseSvc
}

func NewSynkRepository(dbSvc *database.DatabaseSvc) *SynkRepository {
	return &SynkRepository{
		dbSvc: dbSvc,
	}
}

func (sr *SynkRepository) CreateSynk(s *models.Synk) error {
	return sr.dbSvc.DB.Create(s).Error
}

func (sr *SynkRepository) CreateSchema(s *models.Schema) error {
	return sr.dbSvc.DB.Create(s).Error
}

func (sr *SynkRepository) GetSynksFromConnectionId(connectionId uuid.UUID, userId uuid.UUID) []models.Synk {
	synks := make([]models.Synk, 0)
	if err := sr.dbSvc.DB.Preload("connections").Where("connection_id = ? AND user_id = ?", connectionId, userId).Find(&synks).Error; err != nil {
		return nil
	}

	return synks
}

func (sr *SynkRepository) GetSynksForUser(userId uuid.UUID) []models.Synk {
	synks := make([]models.Synk, 0)
	if err := sr.dbSvc.DB.Where("user_id = ?", userId).Find(&synks).Error; err != nil {
		return nil
	}

	return synks
}

func (sr *SynkRepository) GetSynkById(synkId uuid.UUID) *models.Synk {
	synk := &models.Synk{}
	if err := sr.dbSvc.DB.First(synk, synkId).Error; err != nil {
		return nil
	}

	return synk
}

func (sr *SynkRepository) UpdateSynk(s *models.Synk) error {
	return sr.dbSvc.DB.Save(s).Error
}
