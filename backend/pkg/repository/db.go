package repository

import (
	"github.com/vpnvsk/test-tast-fullstack/tree/main/backend/internal/models"
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
