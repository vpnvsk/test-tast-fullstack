package repository

import (
	"github.com/vpnvsk/test-tast-fullstack/tree/main/backend/internal/models"
)

type Users interface {
	CreateUser(user models.User) (int, error)
	GetAllUser() ([]models.User, error)
	GetUserByID(userId int) (models.User, error)
	DeleteUser(userId int) error
	UpdateUser(userId int, input models.UserUpdate) error
}

type Repository struct {
	Users
}

func NewRepository(db *InMemoryDB) *Repository {
	return &Repository{
		Users: NewUserDb(db),
	}
}
