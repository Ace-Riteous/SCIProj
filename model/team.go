package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	StudentID string `gorm:"type:varchar(10);comment:学号" json:"student_id"`
	TeacherID string `gorm:"type:varchar(7);comment:统一认证码" json:"teacher_id"`
	CID       uint   `gorm:"type:int;comment:竞赛号" json:"c_id"`
	LeaderID  string `gorm:"type:varchar(10);comment:队长学号" json:"leader_id"`
	TeamName  string `gorm:"type:varchar(20);not null;comment:团队姓名" json:"team_name"`
	QQGroup   string `gorm:"type:varchar(20);comment:QQ群" json:"qq_group"`
	Now2all   string `gorm:"type:varchar(10);comment:在伍人数" json:"now2all"`
	IsFull    bool   `gorm:"type:bool;comment:是否满员" json:"is_full"`
}
