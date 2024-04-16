package api

import (
	"SCIProj/model"
	"SCIProj/service"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	params := model.GetRequestJsonParam(c)
	newStudent := model.Student{
		params["username"].(string),
		params["password"].(string),
		params["email"].(string),
		params["phone"].(string),
		params["role"].(string),
		params["avatar"].(string),
		params["age"].(int),
		params["studentid"].(string),
		params["my_teacher"].(string),
	}
	err := service.Register(newStudent, c)
	if err != nil {
		model.Error(c, err)
		return
	}
}
