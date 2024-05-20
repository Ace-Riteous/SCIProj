package service

import (
	"SCIProj/dao"
	"SCIProj/dto"
	"SCIProj/model"
	"SCIProj/utils"
	"errors"
	"strconv"
)

var userService *UserService

type UserService struct {
	BaseService
	Dao *dao.UserDao
}

func NewUserService() *UserService {
	if userService == nil {
		userService = &UserService{
			Dao: dao.NewUserDao(),
		}
	}
	return userService
}

func (m UserService) Login(dto dto.UserLoginDTO) (interface{}, error) {
	if err := dto.Validate(); err != nil {
		return "", err
	}
	student, ok := m.Dao.CheckIsStudent(dto)
	if !ok {
		teacher, ok := m.Dao.CheckIsTeacher(dto)
		if !ok {
			return "", errors.New("User not found")
		} else {
			if utils.CompareHashAndPassword(teacher.Password, dto.Password) {
				uid := teacher.SevenID
				token, err := utils.Award(uid)
				if err != nil {
					return "", err
				}
				go UserHealthCheck(teacher.SevenID)
				type reData struct {
					Token string
					Role  string
				}
				return reData{token, "teacher"}, nil
			}
		}
	} else {
		if utils.CompareHashAndPassword(student.Password, dto.Password) {
			uid := student.SevenID
			token, err := utils.Award(uid)
			if err != nil {
				return "", err
			}
			go UserHealthCheck(student.SevenID)
			type reData struct {
				Token string
				Role  string
			}
			return reData{token, "student"}, nil
		}
	}
	return "", errors.New("Password incorrect")
}

func (m UserService) Register(registerDTO dto.UserRegisterDTO) error {
	if err := registerDTO.Validate(); err != nil {
		return err
	}
	_, ok := m.Dao.CheckIsStudent(dto.UserLoginDTO{SevenID: registerDTO.SevenID})
	if ok {
		return errors.New("User already exists")
	}

	registerDTO.Password, _ = utils.Encrypt(registerDTO.Password)
	err := m.Dao.RegisterUser(registerDTO)
	if err != nil {
		return err
	}
	go UserHealthCheck(registerDTO.SevenID)
	return nil
}

func (m UserService) GetUserInfo(claims *utils.Claims) (model.Student, error) {
	uid, err := strconv.Atoi(claims.Uid)
	if err != nil {
		return model.Student{}, err
	}
	student, err := m.Dao.GetStudentBySevenID(int64(uid))
	if err != nil {
		return model.Student{}, err
	}
	return student, nil

}
