package api

import (
	"SCIProj/model"
	"SCIProj/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//接收用户名密码，返回json
	authid := c.PostForm("authid")
	password := c.PostForm("password")
	loginRes, err := service.Login(authid, password, c)
	if err != nil {
		model.Error(c, err)
		return
	}
	model.Success(c, loginRes)
}
