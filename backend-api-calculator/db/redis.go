package db

import (
    "context"
    "log"
    "github.com/redis/go-redis/v9"
)

func ConnectRedis(addr string, password string, db int) *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: password, // leave empty if no password
        DB:       db,       // default DB index
    })

    ctx := context.Background()
    if err := rdb.Ping(ctx).Err(); err != nil {
        log.Fatalf("failed to connect to Redis: %v", err)
    }

    log.Println("Connected to Redis")
    return rdb
}
