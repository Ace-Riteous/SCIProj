package api

import (
	"SCIProj/global"
	"SCIProj/model"
	"SCIProj/service"
	"SCIProj/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

func SeeCompetitions(c *gin.Context) {
	//无需判断登陆状态即可使用
	CompetitionList, err := service.FetchCompetitionList()
	if err != nil {
		model.Error(c, err)
	}
	global.REDIS.Set("CompetitionList", CompetitionList, 0)
	model.Success(c, CompetitionList)
}

func AddCompetition(c *gin.Context) {
	//判断教师登陆状态
	token := c.GetHeader("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		model.Error(c, errors.New("登录已过期！"))
		return
	}
	//获取学生和教师信息，教师信息可以从token中获取，学生信息需要从前端传入
	student, _ := service.GetStudentById(c.PostForm("学生ID"))
	teacher, _ := service.GetTeacherById(claim.Uid)
	var students []string
	for _, s := range student {
		students = append(students, s.Username)
	}
	newArticle := model.Article{
		PId:        c.PostForm("PId"),
		Title:      c.PostForm("比赛名称"),
		Request:    c.PostForm("比赛要求"),
		Content:    c.PostForm("比赛描述"),
		Teacher:    teacher.Username,
		Student:    students,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		DeleteTime: global.EmptyTime,
	}
	err = service.AddCompetition(&newArticle)
	if err != nil {
		model.Error(c, err)
	}
	global.REDIS.Set("NewArticle", newArticle, 0)
	model.Success(c, newArticle)

}
