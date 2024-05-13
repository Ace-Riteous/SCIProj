package api

import (
	"SCIProj/model"
	"SCIProj/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//接收用户名密码
	uid := c.PostForm("sevenid")
	password := c.PostForm("password")
	loginRes, err := service.Login(uid, password)
	if err != nil {
		model.Error(c, err)
		return
	}
	model.Success(c, loginRes)
}
