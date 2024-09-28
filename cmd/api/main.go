package main

import (
	"github.com/devmegablaster/SheetBridge/internal/config"
)

func main() {
	cfg := config.NewConfig()
	cfg.Init()
}
