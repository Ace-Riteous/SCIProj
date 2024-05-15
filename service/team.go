package service

import (
	"SCIProj/dao"
	"SCIProj/dto"
	"SCIProj/model"
	"SCIProj/utils"
	"strconv"
)

var teamService *TeamService

type TeamService struct {
	BaseService
	Dao *dao.TeamDao
}

func NewTeamService() *TeamService {
	if teamService == nil {
		teamService = &TeamService{
			Dao: dao.NewTeamDao(),
		}
	}
	return teamService
}

func (m TeamService) GetTeamAll(dto dto.TeamAllDTO) ([]model.Team, int64, error) {
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
	teamList, nTotal, err := m.Dao.GetTeamAll(page, limit)
	if err != nil {
		return nil, 0, err
	}
	return teamList, nTotal, nil
}

func (m TeamService) GetTeamNotFull(fullDTO dto.TeamAllDTO) ([]model.Team, int64, error) {
	if err := fullDTO.Validate(); err != nil {
		return nil, 0, err
	}
	page, limit := fullDTO.Page, fullDTO.PageSize
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 20 {
		limit = 10
	}
	teamList, nTotal, err := m.Dao.GetTeamNotFull(page, limit)
	if err != nil {
		return nil, 0, err
	}
	return teamList, nTotal, nil

}

func (m TeamService) NewTeam(addDTO dto.TeamAddDTO, token string) error {
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
	return m.Dao.NewTeam(addDTO, strconv.Itoa(uid))
}

func (m TeamService) JoinTeam(joinDTO dto.TeamJoinDTO, token string) error {
	if err := joinDTO.Validate(); err != nil {
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
	return m.Dao.JoinTeam(joinDTO, strconv.Itoa(uid))

}
