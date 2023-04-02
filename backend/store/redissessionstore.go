package store

import (
	"log"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

type RedisSessionStore struct {
	rp *redis.Pool
}

func NewRedisSessionStore(adr string) *RedisSessionStore {
	return &RedisSessionStore{
		rp: &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				conn, err := redis.DialURL(adr)
				if err != nil {
					log.Printf("ERROR: fail init redis pool: %s", err.Error())
					os.Exit(1)
				}
				return conn, err
			},
		}}
}

func (s *RedisSessionStore) Insert(token string, username string) error {
	conn := s.rp.Get()
	defer conn.Close()
	_, err := conn.Do("SETEX", token, "24000", username)

	return err
}

func (s *RedisSessionStore) Get(token string) (string, error) {
	conn := s.rp.Get()
	defer conn.Close()
	res, err := redis.String(conn.Do("GET", token))
	return res, err
}

func (s *RedisSessionStore) Delete(token string) error {
	conn := s.rp.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", token)
	return err
}
