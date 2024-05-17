package api

import (
	"SCIProj/dto"
	"SCIProj/service"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

type PostApi struct {
	BaseApi
	Service *service.PostService
}

func NewPostApi() PostApi {

	return PostApi{
		NewBaseApi(),
		service.NewPostService(),
	}
}

func (m PostApi) GetCompetitionAll(c *gin.Context) {
	var iCompetitionAllDTO dto.CompetitionAllDTO
	if err := m.BuildRequest(BuildRequestOption{
		Ctx:     c,
		DTO:     &iCompetitionAllDTO,
		BindAll: true,
	}).GetError(); err != nil {
		return
	}

	competitionList, nTotal, err := m.Service.GetCompetitionAll(iCompetitionAllDTO)
	if err != nil {
		m.Fail(ResponseJson{
			Code: 10001,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Data:  competitionList,
		Total: nTotal,
	})
}

func (m PostApi) AddCompetition(c *gin.Context) {
	var iCompetitionAddDTO dto.CompetitionAddDTO
	if err := m.BuildRequest(BuildRequestOption{
		Ctx:     c,
		DTO:     &iCompetitionAddDTO,
		BindAll: true,
	}).GetError(); err != nil {
		return
	}
	token := c.GetHeader("Authorization")
	err := m.Service.AddCompetition(iCompetitionAddDTO, token)
	if err != nil {
		m.Fail(ResponseJson{
			Code: 10001,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Msg: "Add competition success",
	})

}

func (m PostApi) GetCompetitionDetailByCid(c *gin.Context) {
	if err := m.BuildRequest(BuildRequestOption{
		Ctx: c,
	}).GetError(); err != nil {
		return
	}
	cidStr := c.Param("c_id")
	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		m.Fail(ResponseJson{
			Code: 10001,
			Msg:  errors.New("cid must be a number").Error(),
		})
		return
	}

	competition, err := m.Service.GetCompetitionDetail(uint(cid))
	if err != nil {
		m.Fail(ResponseJson{
			Code: 10001,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Data: competition,
	})
}
