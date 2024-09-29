package routes

import (
	"github.com/devmegablaster/SheetBridge/api/handlers"
	"github.com/devmegablaster/SheetBridge/api/middlewares"
	"github.com/labstack/echo/v4"
)

func registerConnectionRoutes(g *echo.Group, cfg *RouterConfig) {
	connection := g.Group("/connections")

	authMiddleware := middlewares.NewAuth(cfg.DbSvc, cfg.Cfg)
	g.Use(authMiddleware.Auth)

	connectionHandler := handlers.NewConnectionHandler(cfg.DbSvc, cfg.Cfg)

	connection.POST("", connectionHandler.CreateConnection)
	connection.GET("", connectionHandler.GetConnections)
}
