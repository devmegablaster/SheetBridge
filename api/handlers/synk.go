package handlers

import (
	"net/http"

	"github.com/devmegablaster/SheetBridge/internal/config"
	"github.com/devmegablaster/SheetBridge/internal/database"
	"github.com/devmegablaster/SheetBridge/internal/models"
	"github.com/devmegablaster/SheetBridge/internal/services"
	"github.com/labstack/echo/v4"
)

type SynkHandler struct {
	dbSvc   *database.DatabaseSvc
	cfg     config.Config
	synkSvc *services.SynkService
}

func NewSynkHandler(dbSvc *database.DatabaseSvc, cfg config.Config) *SynkHandler {
	return &SynkHandler{
		dbSvc:   dbSvc,
		cfg:     cfg,
		synkSvc: services.NewSynkService(dbSvc, cfg.Kafka),
	}
}

func (h *SynkHandler) NewSynk(c echo.Context) error {
	synkR := models.SynkRequest{}
	if err := c.Bind(&synkR); err != nil {
		return err
	}

	user := c.Get("user").(*models.User)

	synk, err := h.synkSvc.CreateSynkFromRequest(&synkR, user.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, synk.ToResponse())
}

func (h *SynkHandler) GetSynks(c echo.Context) error {
	user := c.Get("user").(*models.User)

	synks, err := h.synkSvc.GetSynksForUser(user.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, synks)
}
