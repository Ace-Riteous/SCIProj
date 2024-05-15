package dao

import (
	"SCIProj/dto"
	"SCIProj/model"
	"strconv"
)

var userDao *UserDao

type UserDao struct {
	BaseDao
}

// NewUserDao 创建实例和返回
func NewUserDao() *UserDao {
	if userDao == nil {
		userDao = &UserDao{
			NewBaseDao(),
		}
	}
	return userDao
}

func (m UserDao) CheckIsStudent(dto dto.UserLoginDTO) (model.Student, bool) {
	var student model.Student
	err := m.Orm.Model(&model.Student{}).
		Where("seven_id=?", dto.SevenID).
		First(&student).
		Error
	if err != nil {
		return model.Student{}, false
	}
	return student, true
}

func (m UserDao) CheckIsTeacher(dto dto.UserLoginDTO) (model.Teacher, bool) {
	var teacher model.Teacher
	err := m.Orm.Model(&model.Teacher{}).
		Where("seven_id=?", dto.SevenID).
		First(&teacher).
		Error
	if err != nil {
		return model.Teacher{}, false
	}
	return teacher, true
}

func (m UserDao) RegisterUser(registerDTO dto.UserRegisterDTO) error {
	var user model.Student
	registerDTO.ConvertToModel(&user)
	err := m.Orm.Save(&user).Error
	if err != nil {
		return err
	}
	return nil

}

func (m UserDao) GetStudentBySevenID(i int64) (model.Student, error) {
	var student model.Student
	err := m.Orm.Model(&model.Student{}).
		Where("seven_id=?", strconv.FormatInt(i, 10)).
		First(&student).
		Error
	if err != nil {
		return model.Student{}, err
	}
	return student, nil

}
