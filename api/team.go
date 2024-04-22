package api

import (
	"SCIProj/global"
	"SCIProj/model"
	"SCIProj/service"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func NewTeam(c *gin.Context) {
	cid := c.PostForm("cid")
	rand.Seed(time.Now().UnixNano())
	maxid := 1000000000
	minid := 999999999
	var tid int
	for {
		tid = rand.Intn(maxid) + minid
		isExist, err := service.CheckTeamIdExist("team" + strconv.Itoa(tid))
		if err != nil {
			if strings.Compare(err.Error(), "TeamId已存在") == 0 {
				continue
			}
			model.Error(c, err)
			return
		}
		if isExist == false {
			break
		}
	}
	teamid := "team" + strconv.Itoa(tid)
	studentNum, err := service.GetStudentNumsByCid(cid)
	if err != nil {
		model.Error(c, err)
		return
	}
	isfull, err := service.TeamIsFull(studentNum, teamid)
	if err != nil {
		model.Error(c, err)
		return
	}
	newTeam := model.Team{
		Name:       c.PostForm("teamname"),
		TeamId:     teamid,
		StudentIds: c.PostForm("studentids"),
		TeacherId:  c.PostForm("teamteacher"),
		CId:        cid,
		IsFull:     isfull,
	}

	err = service.NewTeam(&newTeam)
	if err != nil {
		model.Error(c, err)
		return

	}
	model.Success(c, newTeam)

}

func JoinTeam(c *gin.Context) {

}

func GetTeamAll(c *gin.Context) {
	//无需判断登陆状态即可使用
	TeamList, err := service.FetchTeamList()
	if err != nil {
		model.Error(c, err)
		return
	}
	global.REDIS.Set("TeamList", TeamList, 0)
	model.Success(c, TeamList)
}

func GetTeamNotFull(c *gin.Context) {

}
