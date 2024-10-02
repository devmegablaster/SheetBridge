package connectors

import (
	"context"
	"fmt"

	"github.com/devmegablaster/SheetBridge/internal/config"
	"github.com/devmegablaster/SheetBridge/internal/database"
	"github.com/devmegablaster/SheetBridge/internal/models"
	"github.com/devmegablaster/SheetBridge/internal/services"
	"github.com/jackc/pgx/v5"
)

type PostgresConnection struct {
	DB             *pgx.Conn
	dbSvc          *database.DatabaseSvc
	conn           *models.Connection
	synk           *models.Synk
	authSvc        *services.AuthService
	userSvc        *services.UserService
	connectionSvc  *services.ConnectionService
	transformerSvc *services.TransformerService
}

func NewPostgresConnection(conn *models.Connection, dbSvc *database.DatabaseSvc, cfg *config.Config, synk ...*models.Synk) (*PostgresConnection, error) {
	connSvc := services.NewConnectionService(dbSvc, cfg.Crypto)
	decryptedConn := connSvc.DecryptConnection(conn)

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		decryptedConn.DatabaseConfig.Host,
		decryptedConn.DatabaseConfig.Port,
		decryptedConn.DatabaseConfig.Username,
		decryptedConn.DatabaseConfig.Password,
		decryptedConn.DatabaseConfig.Database,
	)

	db, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %v", err)
	}

	if len(synk) > 0 {
		return &PostgresConnection{
				DB:            db,
				conn:          conn,
				synk:          synk[0],
				dbSvc:         dbSvc,
				authSvc:       services.NewAuthService(dbSvc, cfg.Crypto, cfg.Auth),
				userSvc:       services.NewUserService(dbSvc),
				connectionSvc: connSvc,
			},
			nil
	}

	return &PostgresConnection{DB: db, conn: conn}, nil
}

func (pc *PostgresConnection) Close() {
	pc.DB.Close(context.Background())
}
