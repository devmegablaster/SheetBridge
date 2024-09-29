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

func NewAuthHandler(dbSvc *database.DatabaseSvc, cfg config.Config) *AuthHandler {
	return &AuthHandler{
		dbSvc:   dbSvc,
		cfg:     cfg,
		authSvc: services.NewAuthService(dbSvc, cfg.Crypto, cfg.Auth),
	}
}

// TODO: Error Handling
func (h *AuthHandler) InitGoogleAuth(c echo.Context) error {
	return h.authSvc.InitGoogleAuth(c.Response(), c.Request())
}

// TODO: Error Handling
func (h *AuthHandler) CallbackGoogleAuth(c echo.Context) error {
	user, err := h.authSvc.CallbackGoogleAuth(c.Response(), c.Request())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

// TODO: Error Handling
func (h *AuthHandler) Login(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	accessToken := token[len("Bearer "):]

	user, err := h.authSvc.Login(accessToken)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}
