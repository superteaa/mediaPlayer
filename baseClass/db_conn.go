package baseClass

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //导入包但不使用，init()
)


func DbConn() (*sql.DB,error){
	db,err := sql.Open("mysql","myweb:520520520@tcp(127.0.0.1:3306)/myweb")

	// db,err := sql.Open("mysql","myweb:520zzj520@tcp(127.0.0.1:3306)/myweb")
	if err != nil {
		fmt.Println("Open mysql failed", err)
		return nil,err
	 }
	return db,nil
}