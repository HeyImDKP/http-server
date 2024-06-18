package main

import (
	"fmt"
	"os"

	lib "bookstore.com/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	mySQLDB := InitDatabase()
	LoadServerFiles(router)
	InitResponse(router, mySQLDB)

	port := os.Getenv("PORT")
	if port != "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		fmt.Println(err)
	}
}

func InitDatabase() *lib.MySQLDB {
	mySQLDB := lib.MySQLDB{
		Username:     os.Getenv("MYSQL_USER"),
		Password:     os.Getenv("MYSQL_PASSWORD"),
		Hostname:     os.Getenv("MYSQL_HOST"),
		DatabaseName: os.Getenv("MYSQL_DATABASE"),
	}
	if err := mySQLDB.ConnectToDB(); err != nil {
		mySQLDB.Username = "davekevin"
		mySQLDB.Password = "pass"
		mySQLDB.Hostname = "localhost:3306"
		mySQLDB.DatabaseName = "users"

		if mySQLDB.ConnectToDB() != nil {
			fmt.Println(err)
		}
	}
	return &mySQLDB
}

func InitResponse(router *gin.Engine, mySQLDB *lib.MySQLDB) {
	router.GET("/", lib.GetLogin())
	router.POST("/login", lib.PostLogin(mySQLDB))
}

func LoadServerFiles(router *gin.Engine) {
	router.StaticFile("/favicon.ico", "assets/favicon.ico")
	router.Static("/assets", "assets")
	router.Static("/styles", "styles")
	router.Static("/scripts", "scripts")
	router.LoadHTMLGlob("template/*")
}
