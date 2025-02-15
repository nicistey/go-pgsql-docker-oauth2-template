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

// // пока не используется
// func SetCache(rdb *redis.Client, key string, value interface{}, ttl time.Duration) error {
//     return rdb.Set(Ctx, key, value, ttl).Err()
// }

// 
// func GetCache(rdb *redis.Client, key string) (string, error) {
//     return rdb.Get(Ctx, key).Result()
// }