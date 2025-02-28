package database

import (
	"sync"
)

// I was forced to create a primitive thread-safe cache. The new bot library requires it.

type Cacher interface {
	SetUser(id int64, user User)
	UpdateUser(user User)
	GetUser(id int64) (User, bool)
	DeleteUser(id int64)

	UpdateUsage(id int64, usage SessionUsage)
	SetUsage(id int64, usage SessionUsage)
	GetUsage(id int64) (SessionUsage, bool)
	DeleteUsage(id int64)
}

type memoryCache struct {
	mu        sync.Mutex
	users     map[int64]User
	usageData map[int64]SessionUsage
}

// Constructor of the Cacher interface
func NewMemoryCache() Cacher {
	return &memoryCache{
		users:     make(map[int64]User),
		usageData: make(map[int64]SessionUsage),
	}
}

func (c *memoryCache) SetUser(id int64, user User) {
	c.mu.Lock()
	c.users[id] = user
	c.mu.Unlock()
}

func (c *memoryCache) UpdateUser(user User) {
	c.mu.Lock()
	c.users[user.ID] = user
	c.mu.Unlock()
}
func (c *memoryCache) GetUser(id int64) (User, bool) {
	c.mu.Lock()
	user, exists := c.users[id]
	c.mu.Unlock()
	return user, exists
}

func (c *memoryCache) DeleteUser(id int64) {
	c.mu.Lock()
	delete(c.users, id)
	c.mu.Unlock()
}
func (c *memoryCache) SetUsage(id int64, usage SessionUsage) {
	c.mu.Lock()
	c.usageData[id] = usage
	c.mu.Unlock()
}
func (c *memoryCache) UpdateUsage(id int64, usage SessionUsage) {
	c.mu.Lock()
	c.usageData[id] = usage
	c.mu.Unlock()
}

func (c *memoryCache) GetUsage(id int64) (SessionUsage, bool) {
	c.mu.Lock()
	usage, exists := c.usageData[id]
	c.mu.Unlock()
	return usage, exists
}

func (c *memoryCache) DeleteUsage(id int64) {
	c.mu.Lock()
	delete(c.usageData, id)
	c.mu.Unlock()
}
