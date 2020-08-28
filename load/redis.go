// Package load redis
package load

import (
	"encoding/json"
	"time"
)
import "github.com/gomodule/redigo/redis"

type redisClient struct {
	*redis.Pool
}

// RedisClient redis客户端
var RedisClient = new(redisClient)

func init() {
	RedisClient.Pool = &redis.Pool{
		MaxIdle:     RedisSetting.MaxIdle,
		MaxActive:   RedisSetting.MaxActive,
		IdleTimeout: RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", RedisSetting.Password); err != nil {
					return nil, c.Close()
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

// Set a key/value
func (r *redisClient) Set(key string, data interface{}, time int) error {
	conn := r.Pool.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

// Get get a key
func (r *redisClient) Get(key string) ([]byte, error) {
	conn := r.Pool.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Delete delete a kye
func (r *redisClient) Del(key string) (bool, error) {
	conn := r.Pool.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}
