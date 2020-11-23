package controller

import (
	"fmt"
	"github.com/Link-lin/BookKeepingBackend/dao"
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
	engine.POST("/api/logout", lc.LogOutUser)
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
	user := us.FindUserByEmail(loginParam.Email)
	if user != nil && user.Loggedin==1{
		tool.Success(context, "User Already Logged In")
		return
	}

	// If user not already logged in
	user = us.Login(loginParam.Email, loginParam.Password)
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

func (lc *LoginController) LogOutUser(context *gin.Context){
	var logoutParam param.LoginParam

	err := tool.Decode(context.Request.Body, &logoutParam)
	if err != nil {
		tool.Failed(context, "Error Parsing the request body")
	}

	us := service.UserService{}

	us.LogOut(logoutParam.Email, logoutParam.Password)
	// There is some error
	ud := dao.UserDao{Orm: tool.DbEngine}
	re := ud.QueryByEmail(logoutParam.Email)

	fmt.Println(re)

	tool.Success(context, re)
}
