package dojwt

import (
	"github.com/golang-jwt/jwt/v5"
    "cloud/baseClass"
    "github.com/redis/go-redis/v9"
    "fmt"
    "context"
)

type Custom struct {
    Username string
    Password string
    Email string
    jwt.RegisteredClaims
}
func Check_token(token string) (string, error) {

    // 检查redis
    key := token
    val, err := baseClass.Client.Get(context.Background(), key).Result()
    if err == redis.Nil {
        fmt.Println("键不存在")
        return "", err
    } else if err != nil {
        fmt.Println("查询错误:", err)
        return "", err
    } else {
        fmt.Println("找到的值:", val)
        return val, nil
    }

    tokenClaims, err := jwt.ParseWithClaims(token, &Custom{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })

    if err != nil {
        return "", err
    }

    if tokenClaims != nil {
        if claims, ok := tokenClaims.Claims.(*Custom); ok && tokenClaims.Valid {
            return claims.Email, nil
        }
    }
    return "", err
}
