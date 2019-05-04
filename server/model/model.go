package model

import (
	"fmt"
	_ "github.com/Go-SQL-Driver/MYSQL"
	"github.com/go-clog/clog"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	mysqlUser := "root"
	mysqlPassword := "huang950413"
	mysqlHost := "127.0.0.1"
	mysqlPort := 3306
	mysqlDatabase := "hc_blog_write"

	address := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4,utf8&parseTime=True", mysqlUser, mysqlPassword,
		mysqlHost, mysqlPort, mysqlDatabase)

	var err error
	db, err = gorm.Open("mysql", address)
	if err != nil {
		clog.Fatal(2, "Failed to open database, err: %v", err)
	}
}
