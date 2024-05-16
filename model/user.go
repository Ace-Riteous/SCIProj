package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Username  string `gorm:"type:varchar(20);not null;comment:学生姓名" json:"username"`
	Password  string `gorm:"type:varchar(100);not null;comment:密码" json:"password"`
	SevenID   string `gorm:"type:varchar(7);not null;comment:统一认证码" json:"seven_id"`
	StudentID string `gorm:"type:varchar(10);comment:学号" json:"student_id"`
}

type Teacher struct {
	gorm.Model
	SevenID  string `gorm:"type:varchar(7);comment:统一认证码" json:"seven_id"`
	Username string `gorm:"type:varchar(20);not null;comment:老师姓名" json:"username"`
	Password string `gorm:"type:varchar(100);not null;comment:密码" json:"password"`
}

type StuTeacher struct {
	gorm.Model
	StudentID string `gorm:"type:varchar(10);comment:学号" json:"student_id"`
	TeacherID string `gorm:"type:varchar(7);comment:统一认证码" json:"teacher_id"`
}
