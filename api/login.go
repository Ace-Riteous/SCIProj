package api

import (
	"SCIProj/model"
	"SCIProj/service"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	//接收用户名密码，返回json
	uid := c.PostForm("uid")
	password := c.PostForm("password")
	err := service.Login(uid, password, c)
	if err != nil {
		model.Error(c, err)
		return
	}
}
