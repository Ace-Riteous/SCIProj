package dto

import (
	"SCIProj/model"
	"SCIProj/utils"
	"github.com/go-playground/validator/v10"
)

type CompetitionAllDTO struct {
	Page     int `json:"page" form:"page" validate:"required,number"`
	PageSize int `json:"page_size" form:"page_size" validate:"required,number"`
}

func (m *CompetitionAllDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

type CompetitionAddDTO struct {
	Title            string `json:"title" form:"title" validate:"required"`
	Request          string `json:"request" form:"request" validate:"required"`
	Member           int    `json:"member" form:"member" validate:"required,number"`
	Content          string `json:"content" form:"content" validate:"required"`
	CompetitionTime  int64  `json:"competition_time" form:"competition_time" validate:"required,number"`
	CompetitionPlace string `json:"competition_place" form:"competition_place" validate:"required"`
	CompetitionLink  string `json:"competition_link" form:"competition_link" validate:"required,url"`
}

func (m *CompetitionAddDTO) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

func (m *CompetitionAddDTO) Convert() model.Competition {
	return model.Competition{
		Title:            m.Title,
		Request:          m.Request,
		Member:           m.Member,
		Content:          m.Content,
		CompetitionTime:  utils.Int64ToTime(m.CompetitionTime),
		CompetitionPlace: m.CompetitionPlace,
		CompetitionLink:  m.CompetitionLink,
	}
}
