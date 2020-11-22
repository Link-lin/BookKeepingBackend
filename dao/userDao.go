package dao

import (
	"fmt"
	"github.com/Link-lin/BookKeepingBackend/model"
	"github.com/Link-lin/BookKeepingBackend/tool"
	"log"
)

type UserDao struct {
	*tool.Orm
}

// Find user with email and password
func (ud *UserDao) Query(email string, password string) *model.User{
	var user model.User

	password = tool.EncoderSha256(password)
	_, err := ud.Where(" user_email = ? and user_passwd = ? ", email, password).Get(&user)
	if err != nil {
		log.Print(err.Error())
		return nil
	}
	return &user
}

// Insert user into Database
func (ud *UserDao) InsertMember(user model.User) int64 {
	result, err := ud.InsertOne(&user)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

// Update LoggedIn Field
func (ud *UserDao) LogOutUser(user model.User) int64{
	re := ud.Query(user.Email, user.Password)
	result, err := ud.Update(re, &model.User{LoggedIn: false})

	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}
