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
	has, err := ud.Where(" email = ? and password = ? ", email, password).Get(&user)
	if !has || err != nil {
		log.Print("Cannot find query")
		return nil
	}

	fmt.Println(user)

	return &user
}

func (ud *UserDao) QueryByEmail(email string) *model.User{
	var user model.User

	has, err := ud.Where("email = ?", email).Get(&user)

	if  !has || err != nil {
		log.Print("Error query or not in query")
		return nil
	}

	return &user
}

// Insert user into Database
func (ud *UserDao) InsertMember(user model.User) int64 {
	fmt.Println("!!!INSERT MEMBER !!!")
	result, err := ud.InsertOne(&user)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}

// Update LoggedIn Field
func (ud *UserDao) LogOutUser(email string) int64{

	// Update only did for things that are not nil or null
	// 这里需要注意，Update会自动从user结构体中提取非0和非nil得值作为需要更新的内容，因此，如果需要更新一个值为0，则此种方法将无法实现：
	result, err := ud.Where("email = ?", email).Update(&model.User{
		Loggedin: -1,
	})

	if err != nil {
		fmt.Println(err.Error())
		return 0
	}

	return result
}
