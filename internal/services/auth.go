package services

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/devmegablaster/SheetBridge/internal/config"
	"github.com/devmegablaster/SheetBridge/internal/database"
	"github.com/devmegablaster/SheetBridge/internal/models"
	"github.com/devmegablaster/SheetBridge/internal/repository"
	"github.com/markbates/goth/gothic"
)

type AuthService struct {
	ur            *repository.UserRepository
	dbSvc         *database.DatabaseSvc
	encryptionSvc *EncryptionService
}

func NewAuthService(dbSvc *database.DatabaseSvc, cfg config.CryptoConfig) *AuthService {
	return &AuthService{
		ur:            repository.NewUserRepository(dbSvc),
		dbSvc:         dbSvc,
		encryptionSvc: NewEncryptionService(cfg),
	}
}

func (s *AuthService) InitGoogleAuth(w http.ResponseWriter, r *http.Request) error {
	gothic.BeginAuthHandler(w, r)
	return nil
}

func (s *AuthService) CallbackGoogleAuth(w http.ResponseWriter, r *http.Request) (models.UserResponse, error) {
	gothUser, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("failed to complete user auth: %w", err)
	}

	dbUser, err := s.ur.GetByEmail(gothUser.Email)
	if err == nil {
		return models.UserResponse{
			Id:    dbUser.Id.String(),
			Email: dbUser.Email,
		}, nil
	}

	encryptedAccessToken, err := s.encryptionSvc.Encrypt(gothUser.AccessToken)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("failed to encrypt access token: %w", err)
	}

	encryptedRefreshToken, err := s.encryptionSvc.Encrypt(gothUser.RefreshToken)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("failed to encrypt refresh token: %w", err)
	}

	newUser := models.User{
		Email:        gothUser.Email,
		AccessToken:  encryptedAccessToken,
		RefreshToken: encryptedRefreshToken,
		ExpiresAt:    gothUser.ExpiresAt,
	}

	if err := s.ur.Create(&newUser); err != nil {
		return models.UserResponse{}, fmt.Errorf("failed to create new user: %w", err)
	}

	slog.Info("Created new user", slog.String("email", newUser.Email))

	return newUser.ToResponse(), nil
}
