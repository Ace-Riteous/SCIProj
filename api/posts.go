package api

import (
	"SCIProj/dto"
	"SCIProj/service"
	"github.com/gin-gonic/gin"
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
