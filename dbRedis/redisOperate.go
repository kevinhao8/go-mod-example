package dbredis

import (
	"fmt"
	"strings"

	"github.com/go-redis/redis"
)

type DBConnetion struct {
	Addr     string
	Password string
	DBIndex  int
	PoolSize int
}

func (db DBConnetion) InitConnetion() (bool, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     db.Addr,
		Password: db.Password, // no password set
		DB:       db.DBIndex,  // use default DB
		PoolSize: db.PoolSize,
	})

	pong, err := client.Ping().Result()
	if err != nil || strings.ToLower(pong) != "pong" {
		fmt.Println(pong, err)
		return false, err
	}
	return true, err
}
