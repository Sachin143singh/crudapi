package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var urIDSN="root:root@tcp(localhost:3306)/mydb?parseTime=true"
var err error
func DataMigration() {
	Database,err=gorm.Open(mysql.Open(urIDSN),&gorm.Config{})
	if err!= nil{
		fmt.Print(err.Error())
		panic("connection faild:(")
	}
	Database.AutoMigrate(&Employee{})

}
