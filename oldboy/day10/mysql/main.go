package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func db *sql.DB // 等一个连接对象

func initDB() (err error) {
	dsn := "root:123456@tcp(192.168.0.151:3306)/gotest"
	// Open 不会校验用户名和密码是否正确，只会校验格式是否正确
	db, err = sql.Open("mysql", dsn) // 这里是等号
	if err != nil {
		// dsn 格式不正确
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func query() {

}

func insert() {

}

func delete() {

}


func main() {
	err := initDB()
	if err != nil {
		fmt.Println("Init DB failed, err: %s\n", err)
		return
	}
	fmt.Println("Connect database success!")
}
