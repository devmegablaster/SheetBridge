package routes

import (
	"github.com/devmegablaster/SheetBridge/internal/config"
	"github.com/devmegablaster/SheetBridge/internal/database"
	"github.com/labstack/echo/v4"
)

type RouterConfig struct {
	DbSvc *database.DatabaseSvc
	Cfg   config.Config
}

func RegisterRoutes(g *echo.Group, cfg *RouterConfig) {
	registerAuthRoutes(g, cfg)
	registerConnectionRoutes(g, cfg)
}
