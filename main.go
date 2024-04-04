package main

import (
    "github.com/gin-gonic/gin"
    "path/filepath"
    "fmt"
    "log"
    "os"
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

    r.LoadHTMLGlob("templates/*")

    // 定义一个路由处理函数来渲染index.html
    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", gin.H{
            "title": "主页",
        })
    })

    // 获取相应的视频资源
    r.GET("/video/:filename", func(c *gin.Context) {

        videoDir := "./videos"
        videoFile := c.Param("filename")
        fullPath := filepath.Join(videoDir, videoFile)
        c.File(fullPath)
    })

    // 获取资源列表
    r.GET("/list", func(c *gin.Context) {

        dirname := "./videos" // 你的音乐文件夹路径
        f, err := os.Open(dirname)
        if err != nil {
            log.Fatal(err)
        }
        defer f.Close()

        files, err := f.Readdirnames(-1) // -1 表示读取所有文件
        if err != nil {
            fmt.Println(err)
            log.Fatal(err)
        }

        c.JSON(200, gin.H{"files":files})
    })

    // 处理multipart表单的POST请求
    r.POST("/upload", func(c *gin.Context) {
        // 从表单中获取文件
        file, err := c.FormFile("file")
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        fmt.Println(file.Filename)

        // dst为文件保存的路径
        dst := "./videos/" + file.Filename
        c.SaveUploadedFile(file, dst)

        // 返回上传成功的消息
        c.JSON(200, gin.H{"message": "文件上传成功"})
    })

    r.Run(":3388") // 在3388端口启动服务
}