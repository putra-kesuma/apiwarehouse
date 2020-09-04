package models

type User struct {
	IdUser string `json:"iduser"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	CreatedAt string `json:"createdat"`
}