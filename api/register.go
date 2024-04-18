package api

import (
	"SCIProj/model"
	"SCIProj/service"
	"SCIProj/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Register(c *gin.Context) {
	age, err := strconv.Atoi(c.PostForm("age"))
	if err != nil {
		model.Error(c, err)
		return
	}
	pwdMd5 := utils.Md5Crypt(c.PostForm("password"))
	newStudent := model.Student{
		Username:  c.PostForm("username"),
		Password:  pwdMd5,
		Email:     c.PostForm("email"),
		Phone:     c.PostForm("phone"),
		Role:      c.PostForm("role"),
		Avatar:    c.PostForm("avatar"),
		Age:       age,
		SevenID:   c.PostForm("sevenid"),
		StudentID: c.PostForm("studentid"),
		MyTeacher: c.PostForm("my_teacher"),
	}
	err = service.Register(newStudent)
	if err != nil {
		model.Error(c, err)
		return
	}
	model.Success(c, nil)
}
