package api

import (
	"SCIProj/global"
	"SCIProj/model"
	"SCIProj/service"
	"SCIProj/utils"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func GetCompetition(c *gin.Context) {
	//无需判断登陆状态即可使用
	CompetitionList, err := service.FetchCompetitionList()
	if err != nil {
		model.Error(c, err)
		return
	}
	global.REDIS.Set("CompetitionList", CompetitionList, 0)
	model.Success(c, CompetitionList)
}

func AddCompetition(c *gin.Context) {
	//获取学生和教师信息，教师信息可以从token中获取，学生信息需要从前端传入
	token := c.GetHeader("Authorization")
	_, claim, _ := utils.ParseToken(token)
	student, _ := service.GetMultyStudentsById(c, c.PostForm("studentid"))
	teacher, _ := service.GetTeacherById(claim.Uid)
	var students string
	for _, s := range student {
		students += s.Username + ","
	}
	cid := time.Now().String()[2:22]
	cid = strings.ReplaceAll(cid, " ", "")
	cid = strings.ReplaceAll(cid, ":", "")
	cid = strings.ReplaceAll(cid, "-", "")
	cid = strings.ReplaceAll(cid, ".", "")
	newCompetition := model.Competition{
		CID:              cid,
		Title:            c.PostForm("title"),
		Request:          c.PostForm("request"),
		Content:          c.PostForm("content"),
		CompetitionTime:  c.PostForm("competitiontime"),
		CompetitionPlace: c.PostForm("competitionplace"),
		CompetitionLink:  c.PostForm("competitionlink"),
		Teacher:          teacher.Username,
		CreateTime:       time.Now(),
		UpdateTime:       time.Now(),
		DeleteTime:       global.EmptyTime,
	}
	err := service.AddCompetition(&newCompetition)
	if err != nil {
		model.Error(c, err)
		return
	}
	global.REDIS.Set("NewCompetition", newCompetition, 0)
	model.Success(c, newCompetition)

}
