package service

import (
	"SCIProj/dao"
	"SCIProj/model"
)

func FetchCompetitionList() (CompetitionList []model.Competition, err error) {
	CompetitionList, err = dao.GetCompetitionAll()
	if err != nil {
		return nil, err
	}
	return CompetitionList, nil
}

func AddCompetition(competition *model.Competition) error {
	err := dao.AddCompetition(competition)
	if err != nil {
		return err
	}
	return nil
}

func CheckCidExist(cid string) (bool, error) {
	isExist, err := dao.CheckCidExist(cid)
	if err != nil {
		return isExist, err
	}
	return isExist, nil
}
