package middlewares

import (
	"net/http"
	"strings"

	"github.com/devmegablaster/SheetBridge/internal/config"
	"github.com/devmegablaster/SheetBridge/internal/database"
	"github.com/devmegablaster/SheetBridge/internal/services"
	"github.com/devmegablaster/SheetBridge/pkg/api"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	dbSvc   *database.DatabaseSvc
	userSvc *services.UserService
	authSvc *services.AuthService
}

func NewAuth(db *database.DatabaseSvc, cfg config.Config) *AuthMiddleware {
	return &AuthMiddleware{
		dbSvc:   db,
		userSvc: services.NewUserService(db),
		authSvc: services.NewAuthService(db, cfg.Crypto, cfg.Auth),
	}
}

func (a *AuthMiddleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bearer := c.Request().Header.Get("Authorization")
		if bearer == "" {
			return c.JSON(http.StatusUnauthorized,
				api.NewResponse(http.StatusUnauthorized, "Unauthorized", nil))
		}

		token := strings.ReplaceAll(bearer, "Bearer ", "")

		id, err := a.authSvc.GetUserIdFromJWT(token)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				api.NewResponse(http.StatusInternalServerError, "Server Error occured, Please try again later!", nil))
		}

		user, err := a.userSvc.GetUserById(uuid.MustParse(id))
		if err != nil {
			return c.JSON(http.StatusUnauthorized,
				api.NewResponse(http.StatusUnauthorized, "Unauthorized", nil))
		}

		c.Set("user", user)

		return next(c)
	}
}
