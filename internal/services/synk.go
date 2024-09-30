package services

import (
	"log/slog"

	"github.com/devmegablaster/SheetBridge/internal/broker"
	"github.com/devmegablaster/SheetBridge/internal/config"
	"github.com/devmegablaster/SheetBridge/internal/database"
	"github.com/devmegablaster/SheetBridge/internal/models"
	"github.com/devmegablaster/SheetBridge/internal/repository"
	"github.com/devmegablaster/SheetBridge/pb"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

type SynkService struct {
	dbSvc        *database.DatabaseSvc
	sr           *repository.SynkRepository
	validator    *validator.Validate
	synkProducer *broker.KafkaProducer
}

func NewSynkService(dbSvc *database.DatabaseSvc, kafkaCfg config.KafkaConfig) *SynkService {
	return &SynkService{
		dbSvc:        dbSvc,
		sr:           repository.NewSynkRepository(dbSvc),
		validator:    validator.New(),
		synkProducer: broker.NewKafkaProducer(kafkaCfg.SynkTopic, kafkaCfg.Partition, kafkaCfg),
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

	protoSynk := pb.Synk{
		Action: pb.Action_INIT,
		Id:     synk.Id.String(),
	}

	bt, err := proto.Marshal(&protoSynk)
	if err != nil {
		slog.Error("Unable to marshal synk")
		return err
	}

	// Produce message to kafka
	s.synkProducer.Produce(bt)
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
