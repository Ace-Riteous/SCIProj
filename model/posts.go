package model

import "time"

type Competition struct {
	CId              string    `gorm:"type:varchar(20);index:cid;primary_key;comment:比赛号" json:"cid"`
	Title            string    `gorm:"type:varchar(100);index:title;comment:比赛标题" json:"title"`
	Request          string    `gorm:"type:longtext;comment:比赛要求" json:"request"`
	Member           int       `gorm:"type:int;not null;comment:团队人数" json:"member"`
	Content          string    `gorm:"type:longtext;comment:比赛内容" json:"content"`
	TeamIDs          string    `gorm:"type:longtext;default:null;comment:团队号" json:"teamid"`
	CompetitionTime  string    `gorm:"type:varchar(50);comment:比赛时间" json:"competitiontime"`
	CompetitionPlace string    `gorm:"type:varchar(50);comment:比赛地点" json:"competitionplace"`
	CompetitionLink  string    `gorm:"type:varchar(50);comment:比赛链接" json:"competitionlink"`
	Teacher          string    `gorm:"type:varchar(50);index:competitionteacher;comment:指导老师" json:"teacher"`
	Student          string    `gorm:"type:varchar(100);index:competitionstudent;comment:学生" json:"student"`
	CreateTime       time.Time `gorm:"autoCreateTime;not null;comment:创建时间" json:"createtime"`
	UpdateTime       time.Time `gorm:"autoUpdateTime;comment:更新时间" json:"updatetime"`
	DeleteTime       time.Time `gorm:"autoDeleteTime;comment:删除时间" json:"deletetime"`
}
