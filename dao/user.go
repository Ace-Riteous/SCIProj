package dao

import (
	"SCIProj/global"
	"SCIProj/model"
	"errors"
)

func GetStudent(uid string, password string) (student *model.Student) {
	global.DB.Table("students").Where("studentid = ? AND password = ?", uid, password).Find(&student)
	return student
}

func GetTeacher(uid string, password string) (teacher *model.Teacher) {
	global.DB.Table("students").Where("teacherid = ? AND password = ?", uid, password).Find(&teacher)
	return teacher
}

func GetTeacherById(id string) (teacher *model.Teacher, err error) {
	global.DB.Table("teachers").Where("id = ?", id).Find(&teacher)
	if teacher.Username == "" {
		return nil, errors.New("没有教师: " + id + " 的信息！")
	}
	return teacher, nil
}

func GetStudentById(id string) (student *model.Student, err error) {
	global.DB.Table("students").Where("id = ?", id).Find(&student)
	if student.Username == "" {
		return nil, errors.New("没有学生: " + id + " 的信息！")
	}
	return student, nil
}

func Register(student model.Student) error {
	global.DB.Table("students").Create(student)
	return nil
}
