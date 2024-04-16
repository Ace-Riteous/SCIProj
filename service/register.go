package service

import (
	"SCIProj/dao"
	"SCIProj/model"
	"errors"
	"github.com/gin-gonic/gin"
)

func Register(student model.Student, c *gin.Context) error {
	if len(student.StudentID) != 7 {
		return errors.New("统一认证码不正确！")
	}
	err := dao.Register(student)
	if err != nil {
		return err
	}
	model.Success(c, nil)
	return nil
}
