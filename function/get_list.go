package function

import (
    "github.com/gin-gonic/gin"
    "fmt"
    "log"
    "os"
	"cloud/dojwt"
)

// 获取资源列表
func Get_list(c *gin.Context) {

	token := c.GetHeader("Authorization")

	email,err := dojwt.Check_token(token)

	if err != nil || email == ""{
		fmt.Println(err)
		c.JSON(200, gin.H{"msg":"token不存在或已过期"})
		return
	}

	dirname := "./videos/" + email // 你的音乐文件夹路径
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
}