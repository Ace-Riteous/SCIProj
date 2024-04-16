package dao

import (
	"SCIProj/global"
	"SCIProj/model"
)

func GetCompetitionAll() (CompetitionList []model.Competition, err error) {
	if err := global.DB.Table("competitions").Model(&model.Competition{}).Find(&CompetitionList).Error; err != nil {
		return nil, err
	}
	if len(CompetitionList) == 0 {
		return nil, nil
	}
	return CompetitionList, nil
}

func AddCompetition(competition *model.Competition) error {
	if err := global.DB.Table("competitions").Create(competition).Error; err != nil {
		return err
	}
	return nil
}
