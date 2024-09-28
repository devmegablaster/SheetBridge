package handlers

import (
	"net/http"

	"github.com/devmegablaster/SheetBridge/internal/config"
	"github.com/devmegablaster/SheetBridge/internal/database"
	"github.com/devmegablaster/SheetBridge/internal/services"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	dbSvc   *database.DatabaseSvc
	cfg     config.Config
	authSvc *services.AuthService
}

// TODO: Error Handling

func NewAuthHandler(dbSvc *database.DatabaseSvc, cfg config.Config) *AuthHandler {
	return &AuthHandler{
		dbSvc:   dbSvc,
		cfg:     cfg,
		authSvc: services.NewAuthService(dbSvc, cfg.Crypto),
	}
}

func (h *AuthHandler) InitGoogleAuth(c echo.Context) error {
	return h.authSvc.InitGoogleAuth(c.Response(), c.Request())
}

func (h *AuthHandler) CallbackGoogleAuth(c echo.Context) error {
	user, err := h.authSvc.CallbackGoogleAuth(c.Response(), c.Request())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
