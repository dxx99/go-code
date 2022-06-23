package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	for i := 0; i < 100000; i++ {
		connect(i)
	}

	time.Sleep(1*time.Minute)
}

func connect(key int)  {
	dsn := "root:123456@tcp(127.0.0.1:13306)/bbsgo_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println(key, "-->", db)
	time.Sleep(1*time.Millisecond)
}



