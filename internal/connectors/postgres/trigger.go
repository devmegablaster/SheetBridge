package connectors

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/devmegablaster/SheetBridge/internal/broker"
	"github.com/devmegablaster/SheetBridge/internal/services"
	"github.com/devmegablaster/SheetBridge/pb"
	"google.golang.org/protobuf/proto"
)

// INFO: Initialize the trigger on table
func (pc *PostgresConnection) InitTrigger() {
	command := fmt.Sprintf(`CREATE OR REPLACE FUNCTION data_change() 
		RETURNS TRIGGER AS $$
		BEGIN
			PERFORM pg_notify('sheetbridge', TG_OP || ' on table %s, row id ' || NEW.id::TEXT);
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;

		CREATE TRIGGER data_trigger
		AFTER INSERT OR UPDATE OR DELETE ON %s
		FOR EACH ROW EXECUTE FUNCTION data_change();`, pc.synk.Table, pc.synk.Table)

	if _, err := pc.DB.Exec(context.Background(), command); err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			slog.Error("Failed to create trigger", slog.String("table", pc.synk.Table), slog.String("error", err.Error()))
		}
	}
}

// TODO: Clean Up
func (uc *PostgresConnection) TriggerRoutine(producer *broker.KafkaProducer) {
	if _, err := uc.DB.Exec(context.Background(), "LISTEN sheetbridge"); err != nil {
		slog.Error("Failed to listen on channel", slog.String("error", err.Error()))
		return
	}

	slog.Info("Trigger Listener initiated", slog.String("userId", uc.conn.UserId.String()), slog.String("synkId", uc.synk.Id.String()))

	for {
		notification, err := uc.DB.WaitForNotification(context.Background())
		if err != nil {
			slog.Error("Error waiting for notification", slog.String("error", err.Error()))
		}
		if notification != nil {
			slog.Info("Notification received", slog.String("payload", notification.Payload))
			user, err := uc.userSvc.GetUserById(uc.conn.UserId)
			if err != nil {
				slog.Error("Unable to get user from connection", slog.String("connectionId", uc.conn.Id.String()))
				continue
			}

			accessToken, err := uc.authSvc.RefreshAccessToken(*user)
			if time.Now().After(user.ExpiresAt) {
				newAccessToken, err := uc.authSvc.RefreshAccessToken(*user)
				if err != nil {
					slog.Error("Unable to refresh access token")
					continue
				}

				accessToken = newAccessToken
			}

			tableData, err := uc.GetTableData(uc.synk.Table)
			transformerSvc := services.NewTransformerService(uc.synk.Schema.Col)
			protoWrite := transformerSvc.TransformToWriteMessage(tableData, pb.WriteType_WRITE_FULL, uc.synk.SpreadsheetId, uc.synk.SheetId, accessToken)
			protoWriteData, err := proto.Marshal(protoWrite)
			producer.Produce(protoWriteData)
		}

		time.Sleep(1 * time.Second)
	}
}
