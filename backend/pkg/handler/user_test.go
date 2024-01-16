package handler

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/vpnvsk/test-tast-fullstack/tree/main/backend/internal/models"
	"github.com/vpnvsk/test-tast-fullstack/tree/main/backend/pkg/repository"
	repositorymocks "github.com/vpnvsk/test-tast-fullstack/tree/main/backend/pkg/repository/mocks"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestHandler_createUser(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *repositorymocks.MockUsers, user models.User)

	tests := []struct {
		name                 string
		inputBody            string
		inputUser            models.User
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"name": "Test Name", "surname": "qwerty", "age": 5}`,
			inputUser: models.User{
				Name:    "Test Name",
				Surname: "qwerty",
				Age:     5,
			},
			mockBehavior: func(r *repositorymocks.MockUsers, user models.User) {
				r.EXPECT().CreateUser(user).Return(1, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"name": "Test Name", "surname": "qwerty", "age": 5}`,
			inputUser: models.User{
				Name:    "Test Name",
				Surname: "qwerty",
				Age:     5,
			},
			mockBehavior: func(r *repositorymocks.MockUsers, user models.User) {
				r.EXPECT().CreateUser(user).Return(0, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			users := repositorymocks.NewMockUsers(c)
			test.mockBehavior(users, test.inputUser)
			repo := &repository.Repository{Users: users}

			handler := Handler{repository: repo}

			// Init Endpoint
			r := gin.New()
			r.POST("/api/user", handler.createUser)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/user",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_getAllUsers(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *repositorymocks.MockUsers)

	tests := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			mockBehavior: func(r *repositorymocks.MockUsers) {
				users := []models.User{
					{Id: 1, Name: "User1", Surname: "Surname1", Age: 25},
					{Id: 2, Name: "User2", Surname: "Surname2", Age: 30},
				}
				r.EXPECT().GetAllUser().Return(users, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `[{"id":1,"name":"User1","surname":"Surname1","age":25},{"id":2,"name":"User2","surname":"Surname2","age":30}]`,
		},
		{
			name: "Service Error",
			mockBehavior: func(r *repositorymocks.MockUsers) {
				r.EXPECT().GetAllUser().Return(nil, errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			users := repositorymocks.NewMockUsers(c)
			test.mockBehavior(users)
			repo := &repository.Repository{Users: users}

			handler := Handler{repository: repo}

			// Init Endpoint
			r := gin.New()
			r.GET("/api/users", handler.getAllUsers)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/users", nil)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
func TestHandler_getUserByID(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *repositorymocks.MockUsers, userID int)

	tests := []struct {
		name                 string
		userID               string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:   "Ok",
			userID: "1",
			mockBehavior: func(r *repositorymocks.MockUsers, userID int) {
				user := models.User{
					Id:      1,
					Name:    "Test User",
					Surname: "UserSurname",
					Age:     25,
				}
				r.EXPECT().GetUserByID(userID).Return(user, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"id":1,"name":"Test User","surname":"UserSurname","age":25}`,
		},
		{
			name:   "Invalid ID",
			userID: "invalid",
			mockBehavior: func(r *repositorymocks.MockUsers, userID int) {
				// This function won't be called in this case
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"strconv.Atoi: parsing \"invalid\": invalid syntax"}`,
		},
		{
			name:   "User Not Found",
			userID: "2",
			mockBehavior: func(r *repositorymocks.MockUsers, userID int) {
				r.EXPECT().GetUserByID(userID).Return(models.User{}, errors.New("user not found"))
			},
			expectedStatusCode:   404,
			expectedResponseBody: `{"message":"user not found"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			users := repositorymocks.NewMockUsers(c)
			userID, _ := strconv.Atoi(test.userID)
			test.mockBehavior(users, userID)
			repo := &repository.Repository{Users: users}

			handler := Handler{repository: repo}

			// Init Endpoint
			r := gin.New()
			r.GET("/api/user/:id", handler.getUserByID)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/api/user/"+test.userID, nil)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_deleteUser(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *repositorymocks.MockUsers, userID int)

	tests := []struct {
		name                 string
		userID               string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:   "Ok",
			userID: "1",
			mockBehavior: func(r *repositorymocks.MockUsers, userID int) {
				r.EXPECT().DeleteUser(userID).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `"ok"`,
		},
		{
			name:   "Invalid ID",
			userID: "invalid",
			mockBehavior: func(r *repositorymocks.MockUsers, userID int) {
				// This function won't be called in this case
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"strconv.Atoi: parsing \"invalid\": invalid syntax"}`,
		},
		{
			name:   "User Not Found",
			userID: "2",
			mockBehavior: func(r *repositorymocks.MockUsers, userID int) {
				r.EXPECT().DeleteUser(userID).Return(errors.New("user not found"))
			},
			expectedStatusCode:   404,
			expectedResponseBody: `{"message":"user not found"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			users := repositorymocks.NewMockUsers(c)
			userID, _ := strconv.Atoi(test.userID)
			test.mockBehavior(users, userID)
			repo := &repository.Repository{Users: users}

			handler := Handler{repository: repo}

			// Init Endpoint
			r := gin.New()
			r.DELETE("/api/user/:id", handler.deleteUser)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("DELETE", "/api/user/"+test.userID, nil)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_updateUser(t *testing.T) {
	// Init Test Table
	type mockBehavior func(r *repositorymocks.MockUsers, userID int, userUpdate models.UserUpdate)

	tests := []struct {
		name                 string
		userID               string
		inputBody            string
		inputUserUpdate      models.UserUpdate
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			userID:    "1",
			inputBody: `{"name": "Updated Name", "surname": "Updated Surname", "age": 30}`,
			inputUserUpdate: models.UserUpdate{
				Name:    getStringPointer("Updated Name"),
				Surname: getStringPointer("Updated Surname"),
				Age:     getIntPointer(30),
			},
			mockBehavior: func(r *repositorymocks.MockUsers, userID int, userUpdate models.UserUpdate) {
				r.EXPECT().UpdateUser(userID, userUpdate).Return(nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `"ok"`,
		},
		{
			name:      "Invalid ID",
			userID:    "invalid",
			inputBody: `{"name": "Updated Name", "surname": "Updated Surname", "age": 30}`,
			mockBehavior: func(r *repositorymocks.MockUsers, userID int, userUpdate models.UserUpdate) {
				// This function won't be called in this case
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"strconv.Atoi: parsing \"invalid\": invalid syntax"}`,
		},
		{
			name:      "User Not Found",
			userID:    "3",
			inputBody: `{"name": "Updated Name", "surname": "Updated Surname", "age": 30}`,
			inputUserUpdate: models.UserUpdate{
				Name:    getStringPointer("Updated Name"),
				Surname: getStringPointer("Updated Surname"),
				Age:     getIntPointer(30),
			},
			mockBehavior: func(r *repositorymocks.MockUsers, userID int, userUpdate models.UserUpdate) {
				r.EXPECT().UpdateUser(userID, userUpdate).Return(errors.New("user not found"))
			},
			expectedStatusCode:   404,
			expectedResponseBody: `{"message":"user not found"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			users := repositorymocks.NewMockUsers(c)
			userID, _ := strconv.Atoi(test.userID)
			test.mockBehavior(users, userID, test.inputUserUpdate)
			repo := &repository.Repository{Users: users}

			handler := Handler{repository: repo}

			// Init Endpoint
			r := gin.New()
			r.PUT("/api/user/:id", handler.updateUser)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", "/api/user/"+test.userID, bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

// Helper functions to convert literals to pointers
func getStringPointer(s string) *string {
	return &s
}

func getIntPointer(i int) *int {
	return &i
}
