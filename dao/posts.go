package dao

import (
	"SCIProj/global"
	"SCIProj/model"
	"errors"
	"gorm.io/gorm"
)

func GetCompetitionAll() (CompetitionList []model.Competition, err error) {
	if err = global.DB.Model(&model.Competition{}).Limit(10).Find(&CompetitionList).Limit(10).Error; err != nil {
		return nil, err
	}
	if len(CompetitionList) == 0 {
		return nil, errors.New("没有查询到比赛信息")
	}
	return CompetitionList, nil
}

func AddCompetition(competition *model.Competition) error {
	if err := global.DB.Create(competition).Error; err != nil {
		return err
	}
	return nil
}

func CheckCidExist(cid int) (bool, error) {
	var competition model.Competition
	err := global.DB.Model(&model.Competition{}).Where("id = ?", cid).First(&competition).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return true, err
	}
	return true, errors.New("CID已存在")

}
