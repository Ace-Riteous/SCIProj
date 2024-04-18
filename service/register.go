package service

import (
	"SCIProj/dao"
	"SCIProj/model"
	"errors"
)

func Register(student model.Student) error {
	if len(student.SevenID) != 7 {
		return errors.New("统一认证码不正确！")
	}
	err := dao.Register(student)
	if err != nil {
		return err
	}
	return nil
}
