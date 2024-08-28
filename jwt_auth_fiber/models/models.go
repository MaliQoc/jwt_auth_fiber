package models

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type LoginResponse struct {
	Token     string `json:"access_token"`
	TokenType string `json:"token_type"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{
	{Username: "admin", Password: "1234"},
}

func IsThereUser(user *User) bool {
	for _, v := range users {
		if v.Username == user.Username && v.Password == user.Password {
			return true
		}
	}
	return false
}

