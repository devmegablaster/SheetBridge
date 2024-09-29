package services

import (
	"github.com/devmegablaster/SheetBridge/internal/database"
	"github.com/devmegablaster/SheetBridge/internal/models"
	"github.com/devmegablaster/SheetBridge/internal/repository"
	"github.com/google/uuid"
)

type UserService struct {
	dbSvc *database.DatabaseSvc
	ur    *repository.UserRepository
}

func NewUserService(dbSvc *database.DatabaseSvc) *UserService {
	return &UserService{
		dbSvc: dbSvc,
		ur:    repository.NewUserRepository(dbSvc),
	}
}

// INFO: Returns user using ID
func (s *UserService) GetUserById(id uuid.UUID) (*models.User, error) {
	return s.ur.GetById(id)
}
