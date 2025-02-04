package redis

import (
	"github.com/redis/go-redis/v9"
)

// RedisConn Redis連線
type RedisConn interface {
	Client() *redis.Client                    // 取得Redis Client
	Lock(name RedisMutexName) error           // 互斥鎖
	Unlock(name RedisMutexName) (bool, error) // 解除互斥鎖
}

// redisConn Redis連線
type redisConn struct {
	client  *redis.Client
	option  *redis.Options
	mutexes map[RedisMutexName]RedisMutex
}

// NewRedisConn 建立Redis連線
func NewRedisConn(fns ...RedisConnOptionFunc) RedisConn {
	client := redis.NewClient(&redis.Options{})
	conn := &redisConn{
		client: client,
		option: &redis.Options{},
	}

	for _, fn := range fns {
		fn(conn)
	}

	return conn
}

// Client 取得Redis Client
func (r *redisConn) Client() *redis.Client {
	return r.client
}

// Lock 互斥鎖
func (r *redisConn) Lock(name RedisMutexName) error {
	mutex, isExist := r.mutexes[name]
	if !isExist {
		panic("mutex not found")
	}

	return mutex.Lock(r.client)
}

// Unlock 解除互斥鎖
func (r *redisConn) Unlock(name RedisMutexName) (bool, error) {
	mutex, isExist := r.mutexes[name]
	if !isExist {
		panic("mutex not found")
	}

	return mutex.Unlock(r.client)
}

// RedisConnOptionFunc Redis連線選項
type RedisConnOptionFunc func(conn *redisConn)

// WithPoolSize 設定連線池大小
func WithPoolSize(poolSize int) RedisConnOptionFunc {
	return func(conn *redisConn) {
		conn.option.PoolSize = poolSize
	}
}

// WithAddr 設定連線地址
func WithAddr(addr string) RedisConnOptionFunc {
	return func(conn *redisConn) {
		conn.option.Addr = addr
	}
}

// WithPassword 設定密碼
func WithPassword(password string) RedisConnOptionFunc {
	return func(conn *redisConn) {
		conn.option.Password = password
	}
}

// WithDB 設定資料庫
func WithDB(db int) RedisConnOptionFunc {
	return func(conn *redisConn) {
		conn.option.DB = db
	}
}

// WithMutex 設定互斥鎖
func WithMutex(name RedisMutexName) RedisConnOptionFunc {
	return func(conn *redisConn) {
		if conn.mutexes == nil {
			conn.mutexes = make(map[RedisMutexName]RedisMutex)
		}

		conn.mutexes[name] = newRedisMutex(name)
	}
}
