package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		var user User

		context.Header("Content-Type", "text/html; charset=utf-8")
		context.BindJSON(&user)
		tokenString, err := GenerateJWT(user)

		if err != nil {
			context.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}

		if valid, err := ValidateJWT(tokenString); !valid || err != nil {
			context.JSON(
				http.StatusBadRequest,
				gin.H{
					"error": err.Error(),
				},
			)
			return
		}

		context.HTML(
			http.StatusOK,
			"template/home.html",
			gin.H{
				"title":   "Book Store",
				"content": "Login Success",
			},
		)
	}
}

func GetLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Content-Type", "text/html; charset=utf-8")
		context.HTML(
			http.StatusOK,
			"template/login.html",
			gin.H{
				"title":   "Book Store",
				"content": "Welcome to Book Store",
			},
		)
	}
}
