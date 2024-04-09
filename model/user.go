package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Student struct {
	User
}

type Teacher struct {
	User
}
