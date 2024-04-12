package baseClass

import (
	"github.com/redis/go-redis/v9"
	"fmt"
	"context"
	"time"
)

var Client *redis.Client

// RedisConn函数创建并返回一个Redis客户端连接
func init() {
    // 初始化Redis客户端
    Client = redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis服务器地址
        Password: "",           // 密码，如果没有则留空
        DB: 0,                  // 数据库编号，默认为0
        ReadTimeout:  time.Second * 3,  // 读取操作的超时时间，这里设置为3秒
        WriteTimeout: time.Second * 3,  // 写入操作的超时时间，这里设置为3秒
        DialTimeout:  time.Second * 5,  // 连接建立的超时时间，这里设置为5秒
    })

    // 检查连接
    pong, err := Client.Ping(context.Background()).Result()
    if err != nil {
        fmt.Println("Redis connection error:", err)
    } else {
        fmt.Println("Redis connected:", pong)
    }
}
