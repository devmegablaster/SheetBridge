package services

import (
	"github.com/devmegablaster/SheetBridge/internal/database"
	"github.com/devmegablaster/SheetBridge/internal/models"
	"github.com/devmegablaster/SheetBridge/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type SynkService struct {
	dbSvc     *database.DatabaseSvc
	sr        *repository.SynkRepository
	validator *validator.Validate
}

func NewSynkService(dbSvc *database.DatabaseSvc) *SynkService {
	return &SynkService{
		dbSvc:     dbSvc,
		sr:        repository.NewSynkRepository(dbSvc),
		validator: validator.New(),
	}
}

// INFO: Create synk from synkRequest
func (s *SynkService) CreateSynkFromRequest(synkR *models.SynkRequest, userId uuid.UUID) (*models.Synk, error) {
	if err := s.validator.Struct(synkR); err != nil {
		return nil, err
	}

	synk := synkR.ToSynk(userId)

	err := s.CreateSynk(synk)
	if err != nil {
		return nil, err
	}

	return synk, nil
}

// INFO: Create a new synk
func (s *SynkService) CreateSynk(synk *models.Synk) error {
	err := s.sr.CreateSynk(synk)
	if err != nil {
		return err
	}

	// TODO: Produce message to kafka
	return nil
}

// INFO: Get synks for user
func (s *SynkService) GetSynksForUser(userId uuid.UUID) ([]models.SynkResponse, error) {
	synks := s.sr.GetSynksForUser(userId)
	if synks == nil {
		return nil, nil
	}

	synkResponse := []models.SynkResponse{}
	for _, s := range synks {
		synkResponse = append(synkResponse, *s.ToResponse())
	}

	return synkResponse, nil
}
