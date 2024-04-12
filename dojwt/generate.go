package dojwt

import (
	"github.com/golang-jwt/jwt/v5"
	"fmt"
    "time"
    "cloud/baseClass"
    "context"
)

var secret string = "Lord my God"

func Generate_token(username string, password string, email string, ch chan<- string){
    claims := Custom{
        Username: username,
        Password: password,
        Email: email,
        RegisteredClaims: jwt.RegisteredClaims {
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 48)), // Token有效期很久
            Issuer: "God",
        },
    }

    token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))

    if err != nil {
        fmt.Println(err)
        ch <- "failure"
        return
    }
    ch <- token // 传入管道

    // 将token存储到Redis
    rdb := baseClass.Client
    err_redis := rdb.Set(context.Background(), token, email, 48*time.Hour).Err()

    if err_redis != nil {
        fmt.Println("存储token到Redis时发生错误:", err_redis)
    }

    fmt.Println("Token已成功存储到Redis")
    return

}