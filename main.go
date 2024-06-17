package main

import (
	"fmt"
	"os"

	response "bookstore.com/lib"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	LoadServerFiles(router)
	InitResponse(router)

	port := os.Getenv("PORT")
	if port != "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		fmt.Println(err)
	}
}

func InitResponse(router *gin.Engine) {
	router.GET("/", response.GetLogin())
	router.POST("/login", response.PostLogin())
}

func LoadServerFiles(router *gin.Engine) {
	router.StaticFile("/favicon.ico", "assets/favicon.ico")
	router.Static("/assets", "assets")
	router.Static("/styles", "styles")
	router.Static("/scripts", "scripts")
	router.LoadHTMLGlob("template/*")
}
