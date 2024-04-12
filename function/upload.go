package function

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"cloud/dojwt"
)

type TokenWithData struct{
	Filename string `json:"filename"`
}

// 处理multipart表单的POST请求
func Upload(c *gin.Context) {

	token := c.GetHeader("Authorization")

	email, err := dojwt.Check_token(token)
	if err != nil || email == ""{
		fmt.Println(err)
		c.JSON(200, gin.H{"msg":"token不存在或已过期"})
		return
	}

	dirname := "./videos/" + email + "/" // 你的音乐文件夹路径

	// 从表单中获取文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}
	fmt.Println(file.Filename)

	// dst为文件保存的路径
	dst := dirname + file.Filename
	c.SaveUploadedFile(file, dst)

	// 返回上传成功的消息
	c.JSON(200, gin.H{"message": "文件上传成功"})
}