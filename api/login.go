package api

import (
	"SCIProj/model"
	"SCIProj/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//接收用户名密码，返回json
	params := model.GetRequestJsonParam(c)
	uid := params["uid"].(string)
	password := params["password"].(string)
	err := service.Login(uid, password, c)
	if err != nil {
		model.Error(c, err)
		return
	}
}
