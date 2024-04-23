package initialize

import (
	"SCIProj/global"
	"SCIProj/model"
	"SCIProj/utils"
	"fmt"
	"time"
)

func TestData() {
	testStudent := model.Student{
		StudentID: "2017212210",
		Password:  utils.Md5Crypt("111111"),
		Username:  "test_student",
		Email:     "1",
		Phone:     "2",
		Role:      "student",
		Avatar:    "3",
		Age:       20,
		SevenID:   "1234567",
		MyTeacher: "test_teacher",
	}
	testTeacher := model.Teacher{
		TeacherID:     "2017212211",
		Password:      utils.Md5Crypt("111111"),
		Username:      "test_teacher",
		Email:         "1",
		Phone:         "2",
		Role:          "teacher",
		Avatar:        "3",
		Age:           20,
		MyStudent:     "test_student",
		MyTeam:        "test_team",
		MyCompetition: "1",
	}

	testTeam := model.Team{
		TeamId:     "team1000000001",
		Name:       "test_team",
		StudentIds: "2017212210,",
		TeacherId:  "2017212211",
		CId:        "1",
		IsFull:     false,
	}
	testCompetition := model.Competition{
		CId:              "c1000000001",
		Title:            "test_title",
		Request:          "test_request",
		Member:           4,
		Content:          "test_content",
		TeamIDs:          "2017212212",
		CompetitionTime:  "2006-01-02 15:04:05",
		CompetitionPlace: "test_place",
		CompetitionLink:  "test_link",
		Teacher:          "test_teacher",
		Student:          "test_student",
		CreateTime:       time.Now(),
		UpdateTime:       time.Now(),
		DeleteTime:       time.Now(),
	}
	global.DB.Unscoped().Delete(&testStudent)
	global.DB.Unscoped().Delete(&testTeacher)
	global.DB.Unscoped().Delete(&testTeam)
	global.DB.Unscoped().Delete(&testCompetition)

	err := global.DB.Model(&model.Student{}).Create(&testStudent).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("pass1")
	err = global.DB.Model(&model.Teacher{}).Create(&testTeacher).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("pass2")
	err = global.DB.Model(&model.Team{}).Create(&testTeam).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("pass3")
	err = global.DB.Model(&model.Competition{}).Create(&testCompetition).Error
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("pass4")
}
