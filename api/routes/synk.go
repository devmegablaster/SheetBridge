package routes

import (
	"github.com/devmegablaster/SheetBridge/api/handlers"
	"github.com/devmegablaster/SheetBridge/api/middlewares"
	"github.com/labstack/echo/v4"
)

func registerSynkRoutes(g *echo.Group, cfg *RouterConfig) {
	synk := g.Group("/synks")

	authMiddleware := middlewares.NewAuth(cfg.DbSvc, cfg.Cfg)
	g.Use(authMiddleware.Auth)

	synkHandler := handlers.NewSynkHandler(cfg.DbSvc, cfg.Cfg)

	synk.POST("", synkHandler.NewSynk)
	synk.GET("", synkHandler.GetSynks)
}
