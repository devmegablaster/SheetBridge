package services

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/devmegablaster/SheetBridge/internal/config"
	"github.com/devmegablaster/SheetBridge/internal/database"
	"github.com/devmegablaster/SheetBridge/internal/models"
	"github.com/devmegablaster/SheetBridge/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/markbates/goth/gothic"
)

type AuthService struct {
	ur            *repository.UserRepository
	dbSvc         *database.DatabaseSvc
	encryptionSvc *EncryptionService
	cfgAuth       config.AuthConfig
}

func NewAuthService(dbSvc *database.DatabaseSvc, cfgCrypto config.CryptoConfig, cfgAuth config.AuthConfig) *AuthService {
	return &AuthService{
		ur:            repository.NewUserRepository(dbSvc),
		dbSvc:         dbSvc,
		encryptionSvc: NewEncryptionService(cfgCrypto),
		cfgAuth:       cfgAuth,
	}
}

// INFO: Initialize Google Auth
func (s *AuthService) InitGoogleAuth(w http.ResponseWriter, r *http.Request) error {
	gothic.BeginAuthHandler(w, r)
	return nil
}

// INFO: Callback for Google Auth
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

// INFO: Get JWT from Google Access Token
func (s *AuthService) Login(accessToken string) (models.UserResponse, error) {
	googleUser, err := s.GetUserFromAccessToken(accessToken)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("failed to get user from access token: %w", err)
	}

	dbUser, err := s.ur.GetByEmail(googleUser.Email)
	if err != nil {
		return models.UserResponse{}, fmt.Errorf("failed to get user by email: %w", err)
	}

	jwt := s.NewJWT(*dbUser)

	responseUser := dbUser.ToResponse()
	responseUser.JWT = jwt

	return responseUser, nil
}

// INFO: Issue new JWT
func (s *AuthService) NewJWT(user models.User) string {
	claims := jwt.MapClaims{
		"id":    user.Id.String(),
		"email": user.Email,
		"iss":   s.cfgAuth.JWTIssuer,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.cfgAuth.JWTSecret))
	if err != nil {
		panic(err)
	}

	return tokenString
}

// INFO: Get claims from JWT
func (s *AuthService) ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfgAuth.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return claims, nil
}

// INFO: Get user ID from JWT
func (s *AuthService) GetUserIdFromJWT(tokenString string) (string, error) {
	claims, err := s.ParseJWT(tokenString)
	if err != nil {
		return "", err
	}

	return claims["id"].(string), nil
}

// INFO: Returns decrypted access token
func (s *AuthService) GetAccessToken(user models.User) string {
	decryptedAccessToken, err := s.encryptionSvc.Decrypt(user.AccessToken)
	if err != nil {
		panic(err)
	}

	return decryptedAccessToken
}

// INFO: Refresh Google access token
func (s *AuthService) RefreshAccessToken(user models.User) (string, error) {
	decryptedRefreshToken, err := s.encryptionSvc.Decrypt(user.RefreshToken)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt refresh token: %w", err)
	}

	resp, err := http.PostForm("https://oauth2.googleapis.com/token", map[string][]string{
		"client_id":     {s.cfgAuth.Google.ClientID},
		"client_secret": {s.cfgAuth.Google.ClientSecret},
		"refresh_token": {decryptedRefreshToken},
		"grant_type":    {"refresh_token"},
	})

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var token map[string]interface{}

	if err = json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return "", err
	}

	encryptedAccessToken, err := s.encryptionSvc.Encrypt(token["access_token"].(string))
	if err != nil {
		return "", fmt.Errorf("failed to encrypt access token: %w", err)
	}

	if err := s.ur.UpdateAccessToken(&user, encryptedAccessToken); err != nil {
		return "", fmt.Errorf("failed to update access token: %w", err)
	}

	slog.Info("Refreshed access token", slog.String("email", user.Email))

	return token["access_token"].(string), nil
}

type GoogleUser struct {
	Email string `json:"email"`
}

// INFO: Get user details from access token
//
// WARN: Scope of the access token should include email
func (s *AuthService) GetUserFromAccessToken(accessToken string) (GoogleUser, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/userinfo?access_token=" + accessToken)

	if err != nil {
		return GoogleUser{}, err
	}

	defer resp.Body.Close()

	var user GoogleUser

	err = json.NewDecoder(resp.Body).Decode(&user)

	if err != nil {
		return GoogleUser{}, err
	}

	return user, nil
}
