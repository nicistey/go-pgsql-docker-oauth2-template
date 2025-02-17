package cache

import (
    "context"

    "github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

// создаем новый клиент для работы с редис
func NewClient(addr, password string, db int) *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     addr,     
        Password: password, 
        DB:       db,	
    })
    return rdb
}