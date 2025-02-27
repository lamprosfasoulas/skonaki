package cache

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
    ctx = context.Background()
    redisClient *redis.Client
)

func InitRedis(){
    redisAddr := os.Getenv("SKON_REDIS_ADDR")
    if redisAddr == "" {
        redisAddr = "localhost:6379"
    }
    redisPasswd := os.Getenv("SKON_REDIS_PASSWD")
    if redisPasswd == "" {
        redisPasswd = ""
    }
    redisDB := 0

    redisClient = redis.NewClient(&redis.Options{
        Addr: redisAddr,
        Password: redisPasswd,
        DB: redisDB,
    })

    if _,err := redisClient.Ping(ctx).Result(); err != nil {
        log.Fatalf("Failed to connect to Redis: %v",err)
    }
}

func GetCont(key string) ([]byte, error) {
    c, e := redisClient.Get(ctx, key).Bytes()
    if e == nil {
        return c, nil
    }
    if e == redis.Nil{
        return nil, nil
    }
    return nil, e
    
}

func SetCont(key string, c []byte) error {
    log.Printf("Storing key: %v into the Cache",key)
    return redisClient.Set(ctx, key, c, 5*time.Hour).Err()
}
