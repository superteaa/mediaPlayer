package function

import (
	"github.com/gin-gonic/gin"
	"cloud/dojwt"
	"fmt"
)

// 获取相应的视频资源
func Get_video(c *gin.Context) {

	var tokenWithData TokenWithData

	err := c.BindJSON(&tokenWithData)

	token := c.GetHeader("Authorization")
	filename := tokenWithData.Filename // 音频名

	email, err := dojwt.Check_token(token)
	if err != nil || email == ""{
		fmt.Println(err)
		c.JSON(200, gin.H{"msg":"token不存在或已过期"})
		return
	}

	videoDir := "./videos/" + email + "/" + filename // 你的音乐文件路径
	c.File(videoDir)
}