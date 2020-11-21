package controller

import "github.com/gin-gonic/gin"

type LoginController struct {

}

func (lc *LoginController) Router (engine *gin.Engine){
	engine.POST("/api/login", lc.LoginUser)
	//engine.GET("/api/login", lc.LoginUser)
	engine.POST("/api/signup", lc.SignUpUser)
}

func (lc *LoginController) LoginUser(context *gin.Context){
	// Get the name and password from context

	context.JSON(200, "Hello")
	// Sha256 encode the password

	// Store the password in database

}

func (lc *LoginController) SignUpUser(context *gin.Context){
	// See if the user email is in db

	// If not create new user and store it in database

}

