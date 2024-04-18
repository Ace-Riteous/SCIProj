package service

import (
	"SCIProj/dao"
	"SCIProj/model"
	"SCIProj/utils"
	"errors"
)

func Login(uid string, password string) (data interface{}, err error) {
	password = utils.Md5Crypt(password)
	student, err := dao.GetStudent(uid, password)
	if err != nil {
		return nil, err
	}
	if student != nil {
		s := model.Student{
			StudentID: student.StudentID,
			Username:  student.Username,
			Email:     student.Email,
			Phone:     student.Phone,
			Role:      student.Role,
			Avatar:    student.Avatar,
			Age:       student.Age,
			SevenID:   student.SevenID,
			MyTeacher: student.MyTeacher,
		}
		slr, err := StudentLogin(s)
		if err != nil {
			return nil, err

		}
		return slr, nil
	}
	teacher, err := dao.GetTeacher(uid, password)
	if err != nil {
		return nil, err
	}
	if teacher != nil {
		t := model.Teacher{
			TeacherID:     teacher.TeacherID,
			Username:      teacher.Username,
			Email:         teacher.Email,
			Phone:         teacher.Phone,
			Role:          teacher.Role,
			Avatar:        teacher.Avatar,
			Age:           teacher.Age,
			MyStudent:     teacher.MyStudent,
			MyTeam:        teacher.MyTeam,
			MyCompetition: teacher.MyCompetition,
		}
		tlr, err := TeacherLogin(t)
		if err != nil {
			return nil, err
		}
		return tlr, nil
	}
	return nil, errors.New("该账号未注册/用户名或密码错误")
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
