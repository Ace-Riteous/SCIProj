package api

import (
	"SCIProj/model"
	"SCIProj/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Register(c *gin.Context) {
	age, err := strconv.Atoi(c.PostForm("age"))
	if err != nil {
		model.Error(c, err)
		return
	}
	newStudent := model.Student{
		Username:  c.PostForm("username"),
		Password:  c.PostForm("password"),
		Email:     c.PostForm("email"),
		Phone:     c.PostForm("phone"),
		Role:      c.PostForm("role"),
		Avatar:    c.PostForm("avatar"),
		Age:       age,
		StudentID: c.PostForm("studentid"),
		MyTeacher: c.PostForm("myteacher"),
	}
	err = service.Register(newStudent, c)
	if err != nil {
		model.Error(c, err)
		return
	}
}
