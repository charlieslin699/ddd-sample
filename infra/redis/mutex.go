package redis

import (
	"sync"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

// RedisMutexName 互斥鎖名稱
type RedisMutexName string

const (
	// RedisMutexNameExample 範例
	RedisMutexNameExample RedisMutexName = "example"
)

// RedisMutex 互斥鎖
type RedisMutex interface {
	Lock(client *redis.Client) error
	Unlock(client *redis.Client) (bool, error)
}

// redisMutex 互斥鎖
type redisMutex func(client *redis.Client) *redsync.Mutex

// newRedisMutex 建立互斥鎖
func newRedisMutex(name RedisMutexName) redisMutex {
	once := new(sync.Once)
	var mutex *redsync.Mutex

	return func(client *redis.Client) *redsync.Mutex {
		once.Do(func() {
			pool := goredis.NewPool(client)
			rs := redsync.New(pool)
			mutex = rs.NewMutex(string(name))
		})

		return mutex
	}
}

// Lock 互斥鎖
func (r redisMutex) Lock(client *redis.Client) error {
	return r(client).Lock()
}

// Unlock 解除互斥鎖
func (r redisMutex) Unlock(client *redis.Client) (bool, error) {
	return r(client).Unlock()
}
