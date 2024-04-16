package model

import "time"

type Article struct {
	PId        string    `json:"id"`
	Title      string    `json:"title"`
	Request    string    `json:"request"`
	Content    string    `json:"content"`
	Teacher    string    `json:"teacher"`
	Student    []string  `json:"student"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	DeleteTime time.Time `json:"delete_time"`
}
