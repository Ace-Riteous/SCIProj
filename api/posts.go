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

func GetCompetition(c *gin.Context) {
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
	student, _ := service.GetMultyStudentsById(c.PostForm("studentid"))
	teacher, _ := service.GetTeacherById(claim.Uid)
	var students []model.Student
	for _, s := range student {
		students = append(students, *s)
	}
	newCompetition := model.Competition{
		CID:     c.PostForm("CID"),
		Title:   c.PostForm("title"),
		Request: c.PostForm("request"),
		Content: c.PostForm("content"),
		Team: model.Team{
			Student: students,
			Teacher: *teacher,
		},
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		DeleteTime: nil,
	}
	err = service.AddCompetition(&newCompetition)
	if err != nil {
		model.Error(c, err)
	}
	global.REDIS.Set("NewCompetition", newCompetition, 0)
	model.Success(c, newCompetition)

}
