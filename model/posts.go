package model

import (
	"gorm.io/gorm"
)

type Competition struct {
	gorm.Model
	Title            string `gorm:"type:varchar(100);comment:比赛标题" json:"title"`
	Request          string `gorm:"type:longtext;comment:比赛要求" json:"request"`
	Member           int    `gorm:"type:int;not null;comment:团队人数" json:"member"`
	Content          string `gorm:"type:longtext;comment:比赛内容" json:"content"`
	TeamIDs          string `gorm:"type:longtext;default:null;comment:团队号" json:"teamid"`
	CompetitionTime  string `gorm:"type:varchar(50);comment:比赛时间" json:"competitiontime"`
	CompetitionPlace string `gorm:"type:varchar(50);comment:比赛地点" json:"competitionplace"`
	CompetitionLink  string `gorm:"type:varchar(50);comment:比赛链接" json:"competitionlink"`
	Teacher          string `gorm:"type:varchar(50);comment:指导老师" json:"teacher"`
	Student          string `gorm:"type:varchar(100);comment:学生" json:"student"`
}
