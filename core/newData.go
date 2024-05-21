package core

import (
	"SCIProj/global"
	"SCIProj/model"
	"SCIProj/utils"
	"errors"
	"gorm.io/gorm"
	"time"
)

func newData() {
	//新建学生
	//新建教师
	//新建比赛
	//新建队伍
	pwd, _ := utils.Encrypt("123456")
	var student = model.Student{
		Username:  "student",
		Password:  pwd,
		SevenID:   "1234567",
		StudentID: "1234567890",
	}
	var teacher = model.Teacher{
		Username: "teacher",
		Password: pwd,
		SevenID:  "7654321",
	}
	var competition = model.Competition{
		Title:            "test",
		Content:          "test",
		Request:          "test",
		Member:           4,
		CompetitionTime:  time.Now(),
		CompetitionLink:  "https://cqupt.edu.cn",
		CompetitionPlace: "cqupt",
	}
	var team = model.Team{
		TeamName:  "test",
		StudentID: "1234567",
		LeaderID:  "1234567",
		CID:       1,
		TeacherID: "7654321",
		QQGroup:   "123456",
		IsFull:    false,
	}
	err := global.DB.Model(&model.Student{}).Where("seven_id=?", student.SevenID).First(&student).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		_ = global.DB.Model(&model.Student{}).Create(&student).Error
	}
	err = global.DB.Model(&model.Teacher{}).Where("seven_id=?", teacher.SevenID).First(&teacher).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		_ = global.DB.Model(&model.Teacher{}).Create(&teacher).Error
	}
	err = global.DB.Model(&model.Competition{}).Where("c_id=?", competition.ID).First(&competition).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		_ = global.DB.Model(&model.Competition{}).Create(&competition).Error
	}
	err = global.DB.Model(&model.Team{}).Where("id=?", team.ID).First(&team).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		_ = global.DB.Model(&model.Competition{}).Create(&competition).Error
	}
}
