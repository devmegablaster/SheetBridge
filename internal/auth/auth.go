package auth

import (
	"github.com/devmegablaster/SheetBridge/internal/config"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func InitGoogle(cfg *config.AuthConfig) {
	goth.UseProviders(
		google.New(
			cfg.Google.ClientID,
			cfg.Google.ClientSecret,
			cfg.Google.CallbackURL,

			// INFO: Scopes
			cfg.Google.Scopes...,
		),
	)
}
