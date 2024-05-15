package model

import (
	"gorm.io/gorm"
	"time"
)

type Competition struct {
	gorm.Model
	Title            string    `gorm:"type:varchar(100);comment:比赛标题" json:"title"`
	Request          string    `gorm:"type:text;comment:比赛要求" json:"request"`
	Member           int       `gorm:"type:int;not null;comment:团队人数" json:"member"`
	Content          string    `gorm:"type:text;comment:比赛内容" json:"content"`
	CompetitionTime  time.Time `gorm:"type:bigint;comment:比赛时间" json:"competition_time"`
	CompetitionPlace string    `gorm:"type:varchar(50);comment:比赛地点" json:"competition_place"`
	CompetitionLink  string    `gorm:"type:varchar(50);comment:比赛链接" json:"competition_link"`
}
