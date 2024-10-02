package processor

import (
	"log/slog"
	"time"

	connectors "github.com/devmegablaster/SheetBridge/internal/connectors/postgres"
	"github.com/devmegablaster/SheetBridge/internal/models"
	"github.com/devmegablaster/SheetBridge/internal/services"
	"github.com/devmegablaster/SheetBridge/pb"
	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
)

func (s *SynkProcessor) init(protoSynk *pb.Synk) {
	synk := s.protoSynkToSynk(protoSynk)
	conn := s.connectionSvc.GetConnectionById(synk.ConnectionId)
	user, err := s.userSvc.GetUserById(synk.UserId)
	if err != nil {
		slog.Error("Unable to get user", slog.String("User", string(synk.UserId.String())))
		return
	}

	dbConn, err := connectors.NewPostgresConnection(conn, s.dbSvc, s.cfg, synk)
	if err != nil {
		slog.Error("Unable to connect to user database", slog.String("User", string(synk.UserId.String())))
		return
	}

	schema := dbConn.GetTableSchema(synk.Table)
	synk.Schema = *schema
	if err := s.sr.UpdateSynk(synk); err != nil {
		slog.Error("Unable to update synk", slog.String("Synk", synk.Id.String()), slog.String("Error", err.Error()))
		return
	}

	slog.Info("Synk Schema Updated", slog.String("Synk", synk.Id.String()))

	tableData, err := dbConn.GetTableData(synk.Table)

	dbConn.InitTrigger()
	go dbConn.TriggerRoutine(s.producer)

	if time.Now().After(user.ExpiresAt) {
		newAccessToken, err := s.authSvc.RefreshAccessToken(*user)
		if err != nil {
			slog.Error("Unable to refresh access token")
			return
		}

		user.AccessToken = newAccessToken
	} else {
		user.AccessToken, err = s.encryptionSvc.Decrypt(user.AccessToken)
		if err != nil {
			slog.Error("Unable to decrypt access token")
			return
		}
	}

	transformerSvc := services.NewTransformerService(synk.Schema.Col)
	pbWrite := transformerSvc.TransformToWriteMessage(tableData, pb.WriteType_WRITE_FULL, synk.SpreadsheetId, synk.SheetId, user.AccessToken)

	bt, err := proto.Marshal(pbWrite)
	if err != nil {
		slog.Error("Unable to marshal write data")
		return
	}

	s.producer.Produce(bt)

	slog.Info("Processed Synk Init", slog.String("Synk", synk.Id.String()))
}

func (s *SynkProcessor) protoSynkToSynk(protoSynk *pb.Synk) *models.Synk {
	return s.sr.GetSynkById(uuid.MustParse(protoSynk.Id))
}
