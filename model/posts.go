package model

import "time"

type Competition struct {
	CID        string    `gorm:"type:varchar(10);index:cid;primary_key;comment:比赛号" json:"cid"`
	Title      string    `gorm:"type:varchar(100);index:title;comment:比赛标题" json:"title"`
	Request    string    `gorm:"type:longtext;comment:比赛要求" json:"request"`
	Content    string    `gorm:"type:longtext;comment:比赛内容" json:"content"`
	Team       string    `gorm:"type:varchar(20);index:competitionteam;comment:团队" json:"team"`
	Teacher    string    `gorm:"type:varchar(50);index:competitionteacher;comment:指导老师" json:"teacher"`
	Student    string    `gorm:"type:varchar(100);index:competitionstudent;comment:学生" json:"student"`
	CreateTime time.Time `gorm:"autoCreateTime;not null;comment:创建时间" json:"createTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime;comment:更新时间" json:"updateTime"`
	DeleteTime time.Time `gorm:"autoDeleteTime;comment:删除时间" json:"deleteTime"`
}
