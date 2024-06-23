package lib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostSignup(mySQLDB *MySQLDB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var (
			userSignup User
		)

		context.Header("Content-Type", "text/html; charset=utf-8")
		err := context.ShouldBind(&userSignup)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !userSignup.Validate() {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user"})
			return
		}

		err = mySQLDB.CreateUser(&userSignup)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.HTML(
			http.StatusOK,
			"template/signup.html",
			gin.H{
				"title":   "Book Store",
				"content": "Signup Success: " + userSignup.Username,
			},
		)
	}
}

func PostLogin(mySQLDB *MySQLDB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var (
			userLogin User
			userDB    User
		)

		context.Header("Content-Type", "text/html; charset=utf-8")
		err := context.ShouldBind(&userLogin)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !userLogin.Validate() {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user"})
			return
		}

		err = mySQLDB.GetUserByUsername(userLogin.Username, &userDB)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if userDB.Password != userLogin.Password {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
			return
		}

		context.HTML(
			http.StatusOK,
			"template/login.html",
			gin.H{
				"title":   "Book Store",
				"content": "Login Success: Logged in as " + userLogin.Username,
			},
		)
	}
}

func GetHome() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header("Content-Type", "text/html; charset=utf-8")
		context.HTML(
			http.StatusOK,
			"template/home.html",
			gin.H{
				"title":   "Book Store",
				"content": "Welcome to Book Store",
			},
		)
	}
}
