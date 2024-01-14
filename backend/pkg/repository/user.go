package repository

import (
	"backend/internal/models"
	"errors"
)

var IndexOutOfRange = errors.New("index out of range")

type UserDB struct {
	db *InMemoryDB
}

func NewUserDb(db *InMemoryDB) *UserDB {
	return &UserDB{db: db}
}

func (u *UserDB) CreateUser(user models.User) (int, error) {
	u.db.mu.Lock()
	defer u.db.mu.Unlock()
	maxID := 0

	for _, user := range u.db.users {
		if user.Id > maxID {
			maxID = user.Id
		}
	}
	user.Id = u.db.users[maxID].Id + 1
	u.db.users[user.Id] = user

	return user.Id, nil
}
func (u *UserDB) GetAllUser() ([]models.User, error) {
	users := make([]models.User, 0, len(u.db.users))
	for _, user := range u.db.users {
		users = append(users, user)
	}

	return users, nil
}
func (u *UserDB) GetUserByID(userId int) (models.User, error) {
	u.db.mu.Lock()
	defer u.db.mu.Unlock()
	user, ok := u.db.users[userId]
	if !ok {
		return user, IndexOutOfRange
	}
	return user, nil
}
func (u *UserDB) DeleteUser(userId int) error {
	u.db.mu.Lock()
	defer u.db.mu.Unlock()
	if _, ok := u.db.users[userId]; ok {
		delete(u.db.users, userId)
		return nil
	}
	return IndexOutOfRange
}
func (u *UserDB) UpdateUser(userId int, input models.UserUpdate) error {
	u.db.mu.Lock()
	defer u.db.mu.Unlock()
	if user, ok := u.db.users[userId]; ok {
		if input.Name != nil {
			user.Name = *input.Name
		}
		if input.Surname != nil {
			user.Surname = *input.Surname
		}
		if input.Age != nil {
			user.Age = *input.Age
		}
		u.db.users[userId] = user

		return nil
	}
	return IndexOutOfRange
}
