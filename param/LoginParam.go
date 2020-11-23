package param

type LoginParam struct {
	Email    string `json:"email"`
	Password string `json:"passwd"`
	LoggedIn int8   `json:"loggedIn"`
}
