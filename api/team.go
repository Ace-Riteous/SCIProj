package api

import (
	"SCIProj/global"
	"SCIProj/model"
	"SCIProj/service"
	"errors"
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
	teamid := c.PostForm("teamid")
	studentid := c.PostForm("studentid")
	team, err := service.FetchTeamByTeamId(teamid)
	if err != nil {
		model.Error(c, err)
		return
	}
	if team.IsFull {
		model.Error(c, errors.New("队伍已满员"))
		return
	}
	studentList := make([]string, 0)
	if team.StudentIds != "" {
		studentList = strings.Split(team.StudentIds, ",")
		for _, v := range studentList {
			if strings.Compare(v, studentid) == 0 {
				model.Error(c, errors.New("学生"+studentid+"已在队伍中"))
				return
			}
		}
		studentList = append(studentList, studentid)
	} else {
		studentList = append(studentList, studentid)
	}
	team.StudentIds = strings.Join(studentList, ",")
	studentNum, err := service.GetStudentNumsByCid(team.CId)
	if err != nil {
		model.Error(c, err)
		return
	}
	isfull, err := service.TeamIsFull(studentNum, teamid)
	if err != nil {
		model.Error(c, err)
		return
	}
	team.IsFull = isfull
	err = service.UpdateTeam(team)
	if err != nil {
		model.Error(c, err)
		return
	}
	model.Success(c, team)
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
	TeamList, err := service.FetchTeamNotFullList()
	if err != nil {
		model.Error(c, err)
		return
	}
	global.REDIS.Set("TeamNotFullList", TeamList, 0)
	model.Success(c, TeamList)
}
