package model

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	UserName string `json:"userName"`
	Address  string `json:"address"`
}
