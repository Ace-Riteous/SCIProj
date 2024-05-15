package service

import (
	"SCIProj/dao"
	"SCIProj/dto"
	"SCIProj/global"
	"SCIProj/model"
	"SCIProj/utils"
	"errors"
	"github.com/spf13/viper"
	"strconv"
	"time"
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

func (m UserService) Login(dto dto.UserLoginDTO) (string, error) {
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
				uid := strconv.Itoa(int(teacher.SevenID))
				token, err := GenerateAndCacheToken(uid)
				if err != nil {
					return "", err
				}
				return token, nil
			}
		}
	} else {
		if utils.CompareHashAndPassword(student.Password, dto.Password) {
			uid := strconv.Itoa(int(student.SevenID))
			token, err := GenerateAndCacheToken(uid)
			if err != nil {
				return "", err
			}
			return token, nil
		}
	}
	return "", errors.New("Password incorrect")
}

func (m UserService) Register(registerDTO dto.UserRegisterDTO) error {
	if err := registerDTO.Validate(); err != nil {
		return err
	}
	err := m.Dao.RegisterUser(registerDTO)
	if err != nil {
		return err
	}
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

func GenerateAndCacheToken(nUserID string) (string, error) {

	token, err := utils.Award(nUserID)
	if err == nil {
		rdKey := utils.GeneralRedisKey(utils.LoginRedisKey, "{id}", nUserID, viper.GetString("custom.prefix"))
	_:
		global.RC.Set(rdKey, token, viper.GetDuration("token.ExpiresDays")*time.Minute*24*60)
	}

	return token, nil

}
