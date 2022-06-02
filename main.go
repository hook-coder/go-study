package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"os"
	"time"
)

type User struct {
	ID        uint
	Name      string
	Age       uint8
	createdAt time.Time
	UpdatedAt time.Time
}

func main() {
	logfile, _ := os.Create("./logs/gin.log")
	gin.DefaultWriter = io.MultiWriter(logfile, os.Stdout)

	db := InitDB()
	defer db.DB()

	//gin.SetMode(gin.ReleaseMode) //解决报错的dbug
	r := gin.Default()
	r.GET("/ping1", func(c *gin.Context) {
		//user := User{ID: 1, Name: "hook", Age: 24}
		//db.Create(user)

		c.JSON(200, gin.H{
			"user":    "user",
			"message": "pong",
		})
	})
	r.Run()
}

func InitDB() *gorm.DB {
	//"user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	//用户名称：用户密码@tcp(地址:端口号)/数据库名称?charset=数据库类型&parsetTime=true
	//parseTime是查询结果是否自动解析为时间。
	//loc是MySQL的时区设置。
	host := "localhost"
	port := "3306"
	database := "blog"
	username := "hook"
	password := "hook"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})

	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	return db
}
