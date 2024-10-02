package processor

import (
	"github.com/devmegablaster/SheetBridge/internal/broker"
	"github.com/devmegablaster/SheetBridge/internal/config"
	"github.com/devmegablaster/SheetBridge/internal/database"
	"github.com/devmegablaster/SheetBridge/internal/repository"
	"github.com/devmegablaster/SheetBridge/internal/services"
	"github.com/devmegablaster/SheetBridge/pb"
)

type SynkProcessor struct {
	dbSvc         *database.DatabaseSvc
	sr            *repository.SynkRepository
	connectionSvc *services.ConnectionService
	authSvc       *services.AuthService
	userSvc       *services.UserService
	encryptionSvc *services.EncryptionService
	cfg           *config.Config
	producer      *broker.KafkaProducer
}

func NewSynkProcessor(dbSvc *database.DatabaseSvc, cfg *config.Config) *SynkProcessor {
	return &SynkProcessor{
		dbSvc:         dbSvc,
		sr:            repository.NewSynkRepository(dbSvc),
		connectionSvc: services.NewConnectionService(dbSvc, cfg.Crypto),
		authSvc:       services.NewAuthService(dbSvc, cfg.Crypto, cfg.Auth),
		userSvc:       services.NewUserService(dbSvc),
		encryptionSvc: services.NewEncryptionService(cfg.Crypto),
		cfg:           cfg,
		producer:      broker.NewKafkaProducer(cfg.Kafka.WriteTopic, cfg.Kafka.Partition, cfg.Kafka),
	}
}

func (s *SynkProcessor) Handle(protoSynk *pb.Synk) {
	switch protoSynk.Action {
	case pb.Action_INIT:
		s.init(protoSynk)
	}
}
