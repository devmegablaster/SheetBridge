package repository

import (
	"github.com/devmegablaster/SheetBridge/internal/database"
	"github.com/devmegablaster/SheetBridge/internal/models"
	"github.com/google/uuid"
)

type UserRepository struct {
	dbSvc *database.DatabaseSvc
}

func NewUserRepository(db *database.DatabaseSvc) *UserRepository {
	return &UserRepository{
		dbSvc: db,
	}
}

func (ur *UserRepository) Create(u *models.User) error {
	return ur.dbSvc.DB.Create(u).Error
}

func (ur *UserRepository) GetByEmail(email string) (*models.User, error) {
	var u models.User
	err := ur.dbSvc.DB.Where("email = ?", email).First(&u).Error
	return &u, err
}

func (ur *UserRepository) GetById(id uuid.UUID) (*models.User, error) {
	var u models.User
	err := ur.dbSvc.DB.Where("id = ?", id).First(&u).Error
	return &u, err
}

func (ur *UserRepository) UpdateAccessToken(u *models.User, accessToken string) error {
	u.AccessToken = accessToken
	return ur.dbSvc.DB.Save(u).Error
}
