package service

import (
	"SCIProj/dao"
	"SCIProj/model"
	"errors"
	"github.com/gin-gonic/gin"
)

func GetTeacherById(teacherId string) (teacher *model.Teacher, err error) {
	teacher, err = dao.GetTeacherById(teacherId)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func GetStudentById(studentId string) (student *model.Student, err error) {
	student, err = dao.GetStudentById(studentId)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func GetMultyStudentsById(c *gin.Context, studentsid ...string) (student []*model.Student, err error) {
	for _, id := range studentsid {
		s, err := dao.GetStudentById(id)
		if err != nil {
			model.Error(c, errors.New("学生: "+id+"不存在"))
			continue
		}
		student = append(student, s)
	}
	return student, nil
}
