package model

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	StudentID int64  `gorm:"type:bigint;comment:学号" json:"studentid"`
	TeacherID int64  `gorm:"type:bigint;comment:统一认证码" json:"teacherid"`
	CID       uint   `gorm:"type:int;comment:竞赛号" json:"cid"`
	Name      string `gorm:"type:varchar(20);not null;comment:团队姓名" json:"teamname"`
	IsFull    bool   `gorm:"type:bool;comment:是否满员" json:"isfull"`
}
