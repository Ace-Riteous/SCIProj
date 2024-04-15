package api

import (
	"SCIProj/global"
	"SCIProj/model"
	"SCIProj/service"
	"github.com/gin-gonic/gin"
	"time"
)

func SeeCompetitions(c *gin.Context) {
	CompetitionList, err := service.FetchCompetitionList()
	if err != nil {
		global.REDIS.Set("CompetitionList", CompetitionList, 0)
		model.Success(c, CompetitionList)
	}
	model.Error(c, err)
}

func AddCompetition(c *gin.Context) {
	newArticle := model.Article{
		PId:        c.PostForm("PId"),
		Title:      c.PostForm("比赛名称"),
		Request:    c.PostForm("比赛要求"),
		Content:    c.PostForm("比赛描述"),
		Teacher:    service.GetTeacherById(c.PostForm("教师ID")),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		DeleteTime: time.Now(),
	}
	//err := c.BindJSON(&newArticle)
	//if err != nil {
	//	model.Error(c, err)
	//}

	err := service.AddCompetition(newArticle)
	if err != nil {
		model.Success(c, newArticle)
	}
	model.Error(c, err)
}
