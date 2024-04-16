package model

type StudentLoginRes struct {
	Token    string  `json:"token"`
	UserInfo Student `json:"userInfo"`
}

type TeacherLoginRes struct {
	Token    string  `json:"token"`
	UserInfo Teacher `json:"userInfo"`
}
