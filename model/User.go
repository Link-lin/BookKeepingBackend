package model

// Structure defined for User
type User struct {
	Email string `xorm:"varchar(30)" json:"user_email"`
	Password string `xorm:"varchar(255)" json:"user_passwd"`
	LoggedIn bool `xorm:"bool" json:"logged_in"`
}
