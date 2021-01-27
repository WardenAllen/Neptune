package database

import (
	"github.com/go-redis/redis"
	log "github.com/golang/glog"
)

var redisdb *redis.Client

func InitRedis() bool {

	//连接服务器
	redisdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	//心跳
	pong, err := redisdb.Ping().Result()
	log.Info(pong, err) // Output: PONG <nil>

	return true
}