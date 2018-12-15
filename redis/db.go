package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var pool *redis.Pool

func InitRedis(addr string, password string, database string) {
	pool = &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,
		Dial : func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if err != nil {
				return  nil, err
			}

			if _, err := conn.Do("AUTH", password); err != nil {
				conn.Close()
				return nil, err
			}

			if _, err := conn.Do("SELECT", database); err != nil {
				conn.Close()
				return nil, err
			}

			return conn, nil
		},
	}
}