package repository

import (
	"github.com/stretchr/testify/assert"
	"github.com/vpnvsk/test-tast-fullstack/tree/main/backend/internal/models"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db := NewInMemoryDB()
	u := NewUserDb(db)
	user1 := models.User{Id: 1, Name: "John"}
	user2 := models.User{Id: 2, Name: "Jane"}
	db.users = map[int]models.User{1: user1, 2: user2}

	newUser := models.User{Id: 0, Name: "NewUser"}
	createdID, err := u.CreateUser(newUser)

	assert.NoError(t, err, "Unexpected error during user creation")

	assert.Greater(t, createdID, 0, "Invalid user ID")

	createdUser, found := db.users[createdID]
	assert.True(t, found, "User not found in the database")

	assert.Equal(t, newUser.Name, createdUser.Name, "User names do not match")
}

func TestGetAllUser(t *testing.T) {
	db := NewInMemoryDB()
	u := NewUserDb(db)
	user1 := models.User{Id: 1, Name: "John"}
	user2 := models.User{Id: 2, Name: "Jane"}
	db.users = map[int]models.User{1: user1, 2: user2}

	users, err := u.GetAllUser()

	assert.NoError(t, err, "Unexpected error while getting all users")

	assert.Len(t, users, len(db.users), "Number of returned users does not match")

	for _, returnedUser := range users {
		_, found := db.users[returnedUser.Id]
		assert.True(t, found, "User not found in the database")
	}

	for _, returnedUser := range users {
		expectedUser := db.users[returnedUser.Id]
		assert.Equal(t, expectedUser, returnedUser, "User data does not match")
	}
}

func TestGetUserByID(t *testing.T) {
	db := NewInMemoryDB()
	u := NewUserDb(db)
	user1 := models.User{Id: 1, Name: "John"}
	user2 := models.User{Id: 2, Name: "Jane"}
	db.users = map[int]models.User{1: user1, 2: user2}

	existingUserID := 1
	existingUser, err := u.GetUserByID(existingUserID)

	assert.NoError(t, err, "Unexpected error while getting an existing user")

	assert.Equal(t, user1, existingUser, "Returned user does not match expected user")

	nonExistingUserID := 999
	_, err = u.GetUserByID(nonExistingUserID)

	assert.EqualError(t, err, IndexOutOfRange.Error(), "Unexpected error for non-existing user ID")
}

func TestDeleteUser(t *testing.T) {
	db := NewInMemoryDB()
	u := NewUserDb(db)
	user1 := models.User{Id: 1, Name: "John"}
	user2 := models.User{Id: 2, Name: "Jane"}
	db.users = map[int]models.User{1: user1, 2: user2}

	existingUserID := 1
	err := u.DeleteUser(existingUserID)

	assert.NoError(t, err, "Unexpected error while deleting an existing user")

	_, found := db.users[existingUserID]
	assert.False(t, found, "User still exists after deletion")

	nonExistingUserID := 999
	err = u.DeleteUser(nonExistingUserID)

	assert.EqualError(t, err, IndexOutOfRange.Error(), "Unexpected error for non-existing user ID")
}

func TestUpdateUser(t *testing.T) {
	db := NewInMemoryDB()
	u := NewUserDb(db)
	user1 := models.User{Id: 1, Name: "John", Surname: "Doe", Age: 25}
	user2 := models.User{Id: 2, Name: "Jane", Surname: "Doe", Age: 30}
	db.users = map[int]models.User{1: user1, 2: user2}

	existingUserID := 1
	validInput := models.UserUpdate{Name: StringPointer("UpdatedName"), Age: IntPointer(26)}
	err := u.UpdateUser(existingUserID, validInput)

	assert.NoError(t, err, "Unexpected error while updating an existing user with valid input")

	updatedUser, found := db.users[existingUserID]
	assert.True(t, found, "User not found after update")
	assert.Equal(t, "UpdatedName", updatedUser.Name, "User name not updated")
	assert.Equal(t, 26, updatedUser.Age, "User age not updated")

	nilInput := models.UserUpdate{}
	err = u.UpdateUser(existingUserID, nilInput)

	assert.NoError(t, err, "Unexpected error while updating an existing user with nil input")

	nonExistingUserID := 999
	err = u.UpdateUser(nonExistingUserID, validInput)

	assert.EqualError(t, err, IndexOutOfRange.Error(), "Unexpected error for non-existing user ID")
}

func StringPointer(s string) *string {
	return &s
}

func IntPointer(i int) *int {
	return &i
}
