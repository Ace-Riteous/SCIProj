package dto

import (
	"SCIProj/model"
	"github.com/go-playground/validator/v10"
)

type TeamAllDTO struct {
	Page     int `json:"page" form:"page" validate:"required,number"`
	PageSize int `json:"page_size" form:"page_size" validate:"required,number"`
}

func (m *TeamAllDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

type TeamAddDTO struct {
	TeamName  string `json:"team_name" form:"team_name" validate:"required"`
	TeacherID string `json:"teacher_id" form:"teacher_id" validate:"required,len=7"`
	CID       uint   `json:"c_id" form:"c_id" validate:"required,number"`
}

func (m *TeamAddDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

func (m *TeamAddDTO) Convert(stuID string, isFull bool) model.Team {
	return model.Team{
		Name:      m.TeamName,
		TeacherID: m.TeacherID,
		CID:       m.CID,
		StudentID: stuID,
		IsFull:    isFull,
	}
}

type TeamJoinDTO struct {
	TeamID int `json:"team_id" form:"team_id" validate:"required,number"`
}

func (m *TeamJoinDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}
