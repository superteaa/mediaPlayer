package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "cloud/function"
)

func main() {

    r := gin.Default()

    r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    })

    fmt.Println("yes")

    api := r.Group("/api")
    {
        api.POST("/logIn", function.Login)
        api.POST("/logUp", function.Logup)
        api.POST("/getList", function.Get_list)
        api.POST("/getVideo", function.Get_video)
        api.POST("/upLoad", function.Upload)
    }

    

    r.Run(":3388") // 在3388端口启动服务
}