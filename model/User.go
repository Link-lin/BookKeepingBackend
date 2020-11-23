package model

// Structure defined for User
type User struct {
	Email string `xorm:"varchar(30)" json:"email"`
	Password string `xorm:"varchar(255)" json:"password"`
	Loggedin bool `xorm:"bool" json:"loggedin"`
}
