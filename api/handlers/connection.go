package handlers

import (
	"net/http"

	"github.com/devmegablaster/SheetBridge/internal/config"
	"github.com/devmegablaster/SheetBridge/internal/database"
	"github.com/devmegablaster/SheetBridge/internal/models"
	"github.com/devmegablaster/SheetBridge/internal/services"
	"github.com/labstack/echo/v4"
)

type ConnectionHandler struct {
	dbSvc         *database.DatabaseSvc
	cfg           config.Config
	connectionSvc *services.ConnectionService
}

func NewConnectionHandler(dbSvc *database.DatabaseSvc, cfg config.Config) *ConnectionHandler {
	return &ConnectionHandler{
		dbSvc:         dbSvc,
		cfg:           cfg,
		connectionSvc: services.NewConnectionService(dbSvc, cfg.Crypto),
	}
}

// TODO: Error Handling
func (h *ConnectionHandler) CreateConnection(c echo.Context) error {
	connectionRequest := models.ConnectionRequest{}
	if err := c.Bind(&connectionRequest); err != nil {
		return err
	}

	user := c.Get("user").(*models.User)
	connection := connectionRequest.ToConnection(user.Id)

	if err := h.connectionSvc.CreateConnection(&connection); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, connection.ToResponse())
}

// TODO: Error Handling
func (h *ConnectionHandler) GetConnections(c echo.Context) error {
	user := c.Get("user").(*models.User)

	connections, err := h.connectionSvc.GetConnectionsForUser(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, connections)
}
