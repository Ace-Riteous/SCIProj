package api

import (
	"SCIProj/dto"
	"SCIProj/service"
	"SCIProj/utils"
	"github.com/gin-gonic/gin"
)

type UserApi struct {
	BaseApi
	Service *service.UserService
}

func NewUserApi() UserApi {

	return UserApi{
		NewBaseApi(),
		service.NewUserService(),
	}
}

func (m *UserApi) Login(c *gin.Context) {
	var iUserLoginDTO dto.UserLoginDTO
	if err := m.BuildRequest(BuildRequestOption{
		Ctx:     c,
		DTO:     &iUserLoginDTO,
		BindAll: true,
	}).GetError(); err != nil {
		return
	}

	token, err := m.Service.Login(iUserLoginDTO)
	if err != nil {
		m.Fail(ResponseJson{
			Code: 10001,
			Msg:  err.Error(),
		})
		return
	}

	type reData struct {
		SevenID string `json:"seven_id"`
		Token   string `json:"token"`
	}
	m.OK(ResponseJson{
		Msg: "Login success",
		Data: reData{
			SevenID: iUserLoginDTO.SevenID,
			Token:   token,
		},
	})

}

func (m *UserApi) Register(c *gin.Context) {
	var iUserRegisterDTO dto.UserRegisterDTO
	if err := m.BuildRequest(BuildRequestOption{
		Ctx:     c,
		DTO:     &iUserRegisterDTO,
		BindAll: true,
	}).GetError(); err != nil {
		return
	}

	err := m.Service.Register(iUserRegisterDTO)
	if err != nil {
		m.Fail(ResponseJson{
			Code: 10001,
			Msg:  err.Error(),
		})
		return
	}
	m.OK(ResponseJson{
		Msg: "Register success",
	})

}

func (m *UserApi) GetUserInfo(c *gin.Context) {
	if err := m.BuildRequest(BuildRequestOption{
		Ctx: c,
	}).GetError(); err != nil {
		return
	}
	t := c.GetHeader("Authorization")
	_, claims, err := utils.ParseToken(t)
	if err != nil {
		m.Fail(ResponseJson{
			Code: 10001,
			Msg:  err.Error(),
		})
		return
	}
	user, err := m.Service.GetUserInfo(claims)
	if err != nil {
		m.Fail(ResponseJson{
			Code: 10001,
			Msg:  err.Error(),
		})
		return
	}

	m.OK(ResponseJson{
		Data: user.SevenID,
	})

}
