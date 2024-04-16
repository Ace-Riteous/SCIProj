package dao

import (
	"SCIProj/global"
	"SCIProj/model"
)

func GetStudent(uid string, password string) (student *model.Student, err error) {
	err = global.DB.Table("students").Where("studentid = ? AND password = ?", uid, password).Find(&student).Error
	if err != nil {
		return nil, err
	}
	return student, nil
}

func GetTeacher(uid string, password string) (teacher *model.Teacher, err error) {
	err = global.DB.Table("students").Where("teacherid = ? AND password = ?", uid, password).Find(&teacher).Error
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func GetTeacherById(id string) (teacher *model.Teacher, err error) {
	err = global.DB.Table("teachers").Where("id = ?", id).Find(&teacher).Error
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func GetStudentById(id string) (student *model.Student, err error) {
	err = global.DB.Table("students").Where("id = ?", id).Find(&student).Error
	if err != nil {
		return nil, err
	}
	return student, nil
}

func Register(student model.Student) error {
	err := global.DB.Table("students").Create(student)
	if err != nil {
		return err.Error
	}
	return nil
}
