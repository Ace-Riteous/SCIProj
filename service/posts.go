package service

import (
	"SCIProj/dao"
	"SCIProj/dto"
	"SCIProj/model"
	"SCIProj/utils"
	"errors"
	"strconv"
)

var postService *PostService

type PostService struct {
	BaseService
	Dao *dao.PostDao
}

func NewPostService() *PostService {
	if postService == nil {
		postService = &PostService{
			Dao: dao.NewPostDao(),
		}
	}
	return postService
}

func (m PostService) GetCompetitionAll(dto dto.CompetitionAllDTO) ([]model.Competition, int64, error) {
	if err := dto.Validate(); err != nil {
		return nil, 0, err
	}
	page, limit := dto.Page, dto.PageSize
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 20 {
		limit = 10
	}
	competitionList, nTotal, err := m.Dao.GetCompetitionAll(page, limit)
	if err != nil {
		return nil, 0, err
	}
	return competitionList, nTotal, nil
}

func (m PostService) AddCompetition(addDTO dto.CompetitionAddDTO, token string) error {
	if err := addDTO.Validate(); err != nil {
		return err
	}
	_, claims, err := utils.ParseToken(token)
	if err != nil {
		return err
	}
	uid, err := strconv.Atoi(claims.Uid)
	if err != nil {
		return err
	}
	if m.Dao.CheckIsTeacher(int64(uid)) {
		return m.Dao.AddCompetition(addDTO)
	} else {
		return errors.New("Only teacher can add competition")
	}
}
