package service

import (
	"SCIProj/dao"
	"SCIProj/model"
	"SCIProj/utils"
	"errors"
	"github.com/gin-gonic/gin"
)

func Login(uid string, password string, c *gin.Context) (data interface{}, err error) {
	password = utils.Md5Crypt(password, "sciproj")
	student, err := dao.GetStudent(uid, password)
	if student == nil {
		teacher, err := dao.GetTeacher(uid, password)
		if teacher == nil {
			return nil, err
		}
		t := model.Teacher{
			TeacherID: teacher.TeacherID,
			Username:  teacher.Username,
			Password:  teacher.Password,
		}
		tlr, err := TeacherLogin(t)
		if err != nil {
			return nil, err
		}
		return tlr, nil
	}
	s := model.Student{
		StudentID: student.StudentID,
		Username:  student.Username,
		Password:  student.Password,
	}
	slr, err := StudentLogin(s)
	if err != nil {
		return nil, err

	}
	return slr, nil
}

func StudentLogin(student model.Student) (*model.StudentLoginRes, error) {
	//生成token，jwt技术
	token, err := utils.Award(student.StudentID)
	if err != nil {
		return nil, errors.New("token未能生成")
	}
	var slr = &model.StudentLoginRes{
		Token:    token,
		UserInfo: student,
	}
	return slr, nil
}

func TeacherLogin(teacher model.Teacher) (*model.TeacherLoginRes, error) {
	token, err := utils.Award(teacher.TeacherID)
	if err != nil {
		return nil, errors.New("token未能生成")
	}
	var tlr = &model.TeacherLoginRes{
		Token:    token,
		UserInfo: teacher,
	}
	return tlr, nil
}
