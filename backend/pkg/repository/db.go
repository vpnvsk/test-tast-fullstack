package repository

import (
	"backend/internal/models"
	"sync"
)

type InMemoryDB struct {
	mu    sync.Mutex
	users map[int]models.User
}

func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		users: make(map[int]models.User),
	}
}
