package dao

import (
	"SCIProj/global"
	"SCIProj/model"
	"errors"
)

func GetCompetitionAll() (CompetitionList []model.Competition, err error) {
	if err := global.DB.Model(&model.Competition{}).Limit(10).Find(&CompetitionList).Error; err != nil {
		return nil, err
	}
	if len(CompetitionList) == 0 {
		return nil, errors.New("没有查询到比赛信息")
	}
	return CompetitionList, nil
}

func AddCompetition(competition *model.Competition) error {
	if err := global.DB.Create(&competition).Error; err != nil {
		return err
	}
	return nil
}
