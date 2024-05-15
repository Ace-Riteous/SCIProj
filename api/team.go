package api

import (
	"SCIProj/dto"
	"SCIProj/service"
	"github.com/gin-gonic/gin"
)

var teamApi *TeamApi

type TeamApi struct {
	BaseApi
	Service *service.TeamService
}

func NewTeamApi() TeamApi {
	return TeamApi{
		NewBaseApi(),
		service.NewTeamService(),
	}
}

func (m TeamApi) GetTeamAll(c *gin.Context) {
	var iTeamAllDTO dto.TeamAllDTO
	if err := m.BuildRequest(BuildRequestOption{
		Ctx:     c,
		DTO:     &iTeamAllDTO,
		BindAll: true,
	}).GetError(); err != nil {
		return
	}

	teamList, nTotal, err := m.Service.GetTeamAll(iTeamAllDTO)
	if err != nil {
		m.Fail(ResponseJson{
			Code: 10001,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Data:  teamList,
		Total: nTotal,
	})
}

func (m TeamApi) GetTeamNotFull(c *gin.Context) {
	var iTeamNotFullDTO dto.TeamAllDTO
	if err := m.BuildRequest(BuildRequestOption{
		Ctx:     c,
		DTO:     &iTeamNotFullDTO,
		BindAll: true,
	}).GetError(); err != nil {
		return
	}

	teamList, nTotal, err := m.Service.GetTeamNotFull(iTeamNotFullDTO)
	if err != nil {
		m.Fail(ResponseJson{
			Code: 10001,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Data:  teamList,
		Total: nTotal,
	})
}

func (m TeamApi) NewTeam(c *gin.Context) {
	var iTeamAddDTO dto.TeamAddDTO
	if err := m.BuildRequest(BuildRequestOption{
		Ctx:     c,
		DTO:     &iTeamAddDTO,
		BindAll: true,
	}).GetError(); err != nil {
		return
	}

	token := c.GetHeader("Authorization")
	err := m.Service.NewTeam(iTeamAddDTO, token)
	if err != nil {
		m.Fail(ResponseJson{
			Code: 10001,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Msg: "new team success",
	})

}

func (m TeamApi) JoinTeam(c *gin.Context) {
	var iTeamJoinDTO dto.TeamJoinDTO
	if err := m.BuildRequest(BuildRequestOption{
		Ctx:     c,
		DTO:     &iTeamJoinDTO,
		BindAll: true,
	}).GetError(); err != nil {
		return
	}

	token := c.GetHeader("Authorization")
	err := m.Service.JoinTeam(iTeamJoinDTO, token)
	if err != nil {
		m.Fail(ResponseJson{
			Code: 10001,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Msg: "join team success",
	})
}
