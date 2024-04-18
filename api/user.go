package api

import (
	"SCIProj/model"
	"SCIProj/service"
	"errors"
	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	//获取用户信息
	StudentID := c.PostForm("studentid")
	TeacherID := c.PostForm("teacherid")
	if StudentID != "" {
		student, err := service.GetStudentById(StudentID)
		if err != nil {
			model.Error(c, err)
		}
		if student == nil {
			model.Error(c, errors.New("学生不存在！"))
		}
		model.Success(c, student)
		return
	}
	if TeacherID != "" {
		teacher, err := service.GetTeacherById(TeacherID)
		if err != nil {
			model.Error(c, err)
		}
		if teacher == nil {
			model.Error(c, errors.New("教师不存在！"))
		}
		model.Success(c, teacher)
		return
	}
	model.Error(c, errors.New("ID不能为空！"))
}
