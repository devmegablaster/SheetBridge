package main

import (
	"fmt"

	"github.com/devmegablaster/SheetBridge/api/routes"
	"github.com/devmegablaster/SheetBridge/internal/auth"
	"github.com/devmegablaster/SheetBridge/internal/config"
	"github.com/devmegablaster/SheetBridge/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.NewConfig()
	cfg.Init()

	db := database.New(&cfg.Database)

	auth.InitGoogle(&cfg.Auth)

	e := echo.New()
	e.Use(middleware.Logger())

	version := e.Group(fmt.Sprintf("/api/%s", cfg.Api.Version))
	routes.RegisterRoutes(version, &routes.RouterConfig{
		DbSvc: db,
		Cfg:   *cfg,
	})

	if err := e.Start(fmt.Sprintf(":%s", cfg.Api.Port)); err != nil {
		panic(err)
	}
}
