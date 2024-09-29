package routes

import (
	"github.com/devmegablaster/SheetBridge/api/handlers"
	"github.com/labstack/echo/v4"
)

func registerAuthRoutes(g *echo.Group, cfg *RouterConfig) {
	auth := g.Group("/auth")

	authHandler := handlers.NewAuthHandler(cfg.DbSvc, cfg.Cfg)

	// INFO: Google Auth
	auth.GET("/google", authHandler.InitGoogleAuth)
	auth.GET("/google/callback", authHandler.CallbackGoogleAuth)

	// INFO: Get JWT
	auth.POST("/login", authHandler.Login)
}
