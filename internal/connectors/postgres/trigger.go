package connectors

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"strings"
	"time"

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
func (uc *PostgresConnection) TriggerRoutine() {
	if _, err := uc.DB.Exec(context.Background(), "LISTEN sheetbridge"); err != nil {
		log.Fatalf("Failed to listen on channel: %v", err)
	}

	fmt.Println("Listening for changes...")

	for {
		notification, err := uc.DB.WaitForNotification(context.Background())
		if err != nil {
			log.Fatalf("Error waiting for notification: %v", err)
		}
		if notification != nil {
			fmt.Println("Received notification:", notification.Payload)
			user, err := uc.userSvc.GetUserById(uc.conn.UserId)
			fmt.Println(uc.conn.UserId)
			if err != nil {
				slog.Error("Unable to get user from connection", slog.String("connectionId", uc.conn.Id.String()))
				fmt.Println(err)
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
			var keyValueList []*pb.KeyValue

			for _, row := range tableData {
				for key, value := range row {
					kv := pb.KeyValue{
						Key:   key,
						Value: fmt.Sprintf("%v", value),
					}
					keyValueList = append(keyValueList, &kv)
				}
			}

			protoWrite := pb.Write{
				WriteType:     pb.WriteType_WRITE_FULL,
				SpreadsheetId: uc.synk.SpreadsheetId,
				SheetName:     uc.synk.SheetId,
				AccessToken:   accessToken,
				WriteData: &pb.WriteData{
					DynamicFields: keyValueList,
				},
			}

			protoWriteData, err := proto.Marshal(&protoWrite)

			uc.producer.Produce(protoWriteData)
		}

		time.Sleep(1 * time.Second)
	}
}
