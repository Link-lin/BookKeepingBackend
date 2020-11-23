package service

import (
	"fmt"
	"github.com/Link-lin/BookKeepingBackend/dao"
	"github.com/Link-lin/BookKeepingBackend/model"
	"github.com/Link-lin/BookKeepingBackend/tool"
)

type UserService struct {
}

func (us *UserService) Login(email string, password string) *model.User {
	ud := dao.UserDao{Orm: tool.DbEngine}
	user := ud.Query(email, password)

	if user != nil {
		return user
	}

	password = tool.EncoderSha256(password)

	user = &model.User{
		Email:    email,
		Password: password,
		Loggedin: 1,
	}

	// Insert user into database
	ud.InsertMember(*user)

	return user
}

func (us *UserService) FindUserByEmail(email string) *model.User{

	ud := dao.UserDao{Orm: tool.DbEngine}
	user := ud.QueryByEmail(email)

	if user == nil{
		return nil
	}

	return user
}

func (us *UserService) LogOut(email string, password string) int64{
	ud := dao.UserDao{Orm: tool.DbEngine}

	/*
	user := model.User{
		Email:    email,
		Password: password,
		LoggedIn: true,
	}
	 */

	result := ud.LogOutUser(email)

	fmt.Println(result)
	fmt.Println(us.FindUserByEmail(email))

	return result
}
