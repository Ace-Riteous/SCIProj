package service

import (
	"SCIProj/dao"
	"SCIProj/model"
	"fmt"
)

func FetchCompetitionList() (CompetitionList []*model.Article, err error) {
	CompetitionList, err = dao.GetCompetitionAll()
	if err != nil {
		return nil, err
	}
	return CompetitionList, nil
}

func AddCompetition(article *model.Article) error {
	err := dao.AddCompetition(article)
	if err != nil {
		return err
	}
	return nil
}

func GetTeacherById(teacherId string) (teacher *model.Teacher, err error) {
	teacher, err = dao.GetTeacherById(teacherId)
	if err != nil {
		return nil, err
	}
	return teacher, nil
}

func GetStudentById(studentsid ...string) (student []*model.Student, err error) {
	for _, id := range studentsid {
		s, err := dao.GetStudentById(id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		student = append(student, s)
	}
	return student, nil
}
