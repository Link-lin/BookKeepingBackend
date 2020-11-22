package controller

import (
	"github.com/Link-lin/BookKeepingBackend/param"
	"github.com/Link-lin/BookKeepingBackend/service"
	"github.com/Link-lin/BookKeepingBackend/tool"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
}

func (lc *LoginController) Router(engine *gin.Engine) {
	engine.POST("/api/login", lc.LoginUser)
	//engine.GET("/api/login", lc.LoginUser)
	engine.POST("/api/signup", lc.SignUpUser)
}

func (lc *LoginController) LoginUser(context *gin.Context) {
	// Get the name and password from context
	var loginParam param.LoginParam

	err := tool.Decode(context.Request.Body, &loginParam)

	if err != nil {
		tool.Failed(context, "Error Parsing the request body")
	}

	// Store the password in database
	us := service.UserService{}
	user := us.Login(loginParam.Email, loginParam.Password)

	if user != nil {
		tool.Success(context, user)
		return
	}
	tool.Failed(context, "Failed to Login")
}

func (lc *LoginController) SignUpUser(context *gin.Context) {
	// See if the user email is in db

	// If not create new user and store it in database

}
