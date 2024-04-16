package function

import (
	"github.com/gin-gonic/gin"
	"cloud/dojwt"
	"fmt"
	"os"
)

func Delete_video(c *gin.Context) {

	var tokenWithData TokenWithData

	err := c.BindJSON(&tokenWithData)

	token := c.GetHeader("Authorization")
	filename := tokenWithData.Filename // 视频文件名

	email, err := dojwt.Check_token(token)
	if err != nil || email == ""{
		fmt.Println(err)
		c.JSON(200, gin.H{"msg":"token不存在或已过期"})
		return
	}

	videoDir := "./videos/" + email + "/" + filename // 视频文件路径

	err = os.Remove(videoDir)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"msg": "删除视频文件失败"})
		return
	}

	c.JSON(200, gin.H{"msg": "删除成功"})
}
