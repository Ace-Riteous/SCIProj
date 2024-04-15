package model

type User struct {
	ID       string `json:"id"`
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
	SID       string `json:"sid"`
	School    string `json:"school"`
	MyTeacher string `json:"my_teacher"`
}

type Teacher struct {
	User
	TeacherID string `json:"teacherid"`
	School    string `json:"school"`
}

type Team struct {
	Student []Student `json:"teamstudent"`
	Teacher Teacher   `json:"teamteacher"`
	TeamId  string    `json:"teamid"`
	Name    string    `json:"teamname"`
}
