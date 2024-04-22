package dao

import (
	"SCIProj/global"
	"SCIProj/model"
	"errors"
)

func GetStudent(uid string, password string) (student *model.Student, err error) {
	err = global.DB.Model(model.Student{}).Where("student_id = ? AND password = ?", uid, password).First(&student).Error
	if err != nil {
		return nil, err
	}
	return student, nil
}

func GetTeacher(uid string, password string) (teacher *model.Teacher, err error) {
	err = global.DB.Model(model.Teacher{}).Where("teacher_id = ? AND password = ?", uid, password).First(&teacher).Error
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func GetTeacherById(id string) (teacher *model.Teacher, err error) {
	err = global.DB.Model(model.Teacher{}).Where("teacher_id = ?", id).First(&teacher).Error
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func GetStudentById(id string) (student *model.Student, err error) {
	err = global.DB.Model(model.Student{}).Where("student_id = ?", id).First(&student).Error
	if err != nil {
		return nil, err
	}
	return student, nil
}

func Register(student model.Student) error {
	err := global.DB.Model(model.Student{}).Create(&student).Error
	if err != nil {
		return errors.New("注册失败")
	}
	return nil
}

func GetStudentNumsByCid(cid string) (num int, err error) {
	var competition model.Competition
	err = global.DB.Model(&model.Competition{}).Where("cid = ?", cid).First(&competition).Error
	if err != nil {
		return 0, err
	}
	num = competition.Member
	return num, nil
}
