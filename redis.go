package monitor

import (
	"time"

	"github.com/go-redis/redis"
)

// RedisMonitor is the type for monitoring redis servers
type RedisMonitor struct {
	BaseMonitor
}

// NewRedisMonitor returns an instance of RedisMonitor
func NewRedisMonitor(name, uri string, frequency time.Duration, alertFunc func() error) *RedisMonitor {
	r := &RedisMonitor{}
	r.Name = name
	r.URI = uri
	r.Frequency = frequency
	r.AlertFunc = alertFunc
	return r
}

// Ping pings the redis
func (r *RedisMonitor) Ping() bool {
	rdb := redis.NewClient(&redis.Options{
		Addr: r.URI,
	})
	if _, err := rdb.Ping().Result(); err != nil {
		return false
	}
	return true
}
