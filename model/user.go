package model

type Student struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Avatar    string `json:"avatar"`
	Age       int    `json:"age"`
	StudentID string `json:"studentid"`
	MyTeacher string `json:"my_teacher"`
}

type Teacher struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Avatar    string `json:"avatar"`
	Age       int    `json:"age"`
	TeacherID string `json:"teacherid"`
}

type Team struct {
	Student []Student `json:"teamstudent"`
	Teacher Teacher   `json:"teamteacher"`
	TeamId  string    `json:"teamid"`
	Name    string    `json:"teamname"`
}
