package processor

import (
	"log/slog"

	connectors "github.com/devmegablaster/SheetBridge/internal/connectors/postgres"
	"github.com/devmegablaster/SheetBridge/internal/models"
	"github.com/devmegablaster/SheetBridge/pb"
	"github.com/google/uuid"
)

func (s *SynkProcessor) init(protoSynk *pb.Synk) {
	synk := s.protoSynkToSynk(protoSynk)
	conn := s.connectionSvc.GetConnectionById(synk.ConnectionId)

	dbConn, err := connectors.NewPostgresConnection(conn, s.dbSvc, s.cfg, synk)
	if err != nil {
		slog.Error("Unable to connect to user database", slog.String("User", string(synk.UserId.String())))
		return
	}

	dbConn.InitTrigger()
	go dbConn.TriggerRoutine()
}

func (s *SynkProcessor) protoSynkToSynk(protoSynk *pb.Synk) *models.Synk {
	return s.sr.GetSynkById(uuid.MustParse(protoSynk.Id))
}
