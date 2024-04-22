package model

type StudentLoginRes struct {
	Token    string  `json:"token"`
	UserInfo Student `json:"userinfo"`
}

type TeacherLoginRes struct {
	Token    string  `json:"token"`
	UserInfo Teacher `json:"userinfo"`
}
