package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Username  string `gorm:"type:varchar(20);not null;comment:学生姓名" json:"username"`
	Password  string `gorm:"type:varchar(100);not null;comment:密码" json:"password"`
	SevenID   int64  `gorm:"type:bigint;not null;comment:统一认证码" json:"seven_id"`
	StudentID int64  `gorm:"type:bigint;comment:学号" json:"student_id"`
}

type Teacher struct {
	gorm.Model
	SevenID  int64  `gorm:"type:bigint;comment:统一认证码" json:"seven_id"`
	Username string `gorm:"type:varchar(20);not null;comment:老师姓名" json:"username"`
	Password string `gorm:"type:varchar(100);not null;comment:密码" json:"password"`
}

type StuTeacher struct {
	gorm.Model
	StudentID  int64 `gorm:"type:bigint;comment:学号" json:"student_id"`
	TeacherID  int64 `gorm:"type:bigint;comment:统一认证码" json:"teacher_id"`
	StuSevenID int64 `gorm:"type:bigint;comment:统一认证码" json:"stu_seven_id"`
}

type StuCompetition struct {
	gorm.Model
	StudentID int64 `gorm:"type:bigint;comment:学号" json:"student_id"`
	CID       uint  `gorm:"type:int;comment:竞赛号" json:"c_id"`
	SevenID   int64 `gorm:"type:bigint;comment:统一认证码" json:"seven_id"`
}

type TeaCompetition struct {
	gorm.Model
	TeacherID int64 `gorm:"type:int64;comment:统一认证码" json:"teacher_id"`
	CID       uint  `gorm:"type:int;comment:竞赛号" json:"c_id"`
}
