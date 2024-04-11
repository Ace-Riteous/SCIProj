package model

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Avatar   string `json:"avatar"`
	Age      int    `json:"age"`
}

type Student struct {
	User
	SID       int    `json:"sid"`
	School    string `json:"school"`
	MyTeacher string `json:"my_teacher"`
}

type Teacher struct {
	User
	TID    int    `json:"tid"`
	School string `json:"school"`
}
