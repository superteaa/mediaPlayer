package function

import (
	"github.com/gin-gonic/gin"
	"cloud/dojwt"
	"cloud/baseClass"
	"fmt"
	"os"
)

var body struct{
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}

// login 处理登录请求
func Login(c *gin.Context) {


    if err := c.BindJSON(&body); err != nil {
        c.JSON(200, gin.H{
			"msg": "data error",
		})	
        return
    }
    if body.Username == "" || body.Password == "" {
        c.JSON(200, gin.H{
			"msg": "empty enter",
		})		
        return
    }

	username := body.Username
	password := body.Password

    // 添加处理用户登录逻辑，比如检查数据库中的用户，生成token等
    db,_ := baseClass.DbConn()

    rows,err := db.Query("select email from user where username = ? and password = ?", username, password)

    if err != nil {
		c.JSON(200, gin.H {
			"msg": "database error",
		})
		db.Close()
    	return
    } else if !rows.Next() {
		c.JSON(200, gin.H {
			"msg": "no this accout",
		})
		db.Close()
        return
    } else {
        // 用户存在，返回token
		var email string
		rows.Scan(&email)

		login_ch := make(chan string)

		go dojwt.Generate_token(username, password, email, login_ch)

		token := <-login_ch

		c.JSON(200, gin.H{
			"msg": "yeah",
			"token":token,
		})
		db.Close()

		return
    }
}

// logup 处理登录请求
func Logup(c *gin.Context) {

    // 解析前端发送的账号密码
	if err := c.BindJSON(&body); err != nil {
        c.JSON(200, gin.H{
			"msg": "data error",
		})	
        return
    }
	username := body.Username
	password := body.Password
	email := body.Email

    if username == "" || password == "" || email == ""{
        c.JSON(200, gin.H{
			"msg":"empty enter",
		})
        return
    }

    // 添加处理用户登录逻辑，比如检查数据库中的用户，生成token等
    db,_ := baseClass.DbConn()

	rows,err := db.Query("select username from user where email = ?", email)

    if err != nil {
		fmt.Println("1")
		c.JSON(200, gin.H{
			"msg":"database error",
		})
        return
    } else if !rows.Next() {

		sqlStr := `INSERT INTO user(username, password, email) VALUES (?, ?, ?)`
		value := [3]string{username, password, email}

		// 执行 SQL 语句
		_, err_add := db.Exec(sqlStr, value[0], value[1], value[2])
		if err_add != nil {

			c.JSON(200, gin.H{
				"msg":"database error",
			})
			fmt.Println(err_add)
			db.Close()
			return

		} else {

			logup_ch := make(chan string)
			go dojwt.Generate_token(username, password, email, logup_ch)
		
			token := <-logup_ch
		
			if token == "failure" {
				c.JSON(500, gin.H{
					"msg": "Internal Server Error",
				})
				return
			}
		
			c.JSON(200, gin.H{
				"msg": "yeah",
				"token": token,
			})
			db.Close()

			// 创建一个新的文件夹
			dirname := "./videos/" + email
			err_file := os.Mkdir(dirname, 0755)
			if err_file != nil {
				c.JSON(200, gin.H{
					"msg":"user already existed",
				})
				fmt.Println(err_file)
			}
			return
		}
	} else {
		c.JSON(200, gin.H{
			"msg":"user already existed",
		})
		return
	}
}