package dao

import (
	"SCIProj/global"
	"SCIProj/model"
)

func GetCompetitionAll() (CompetitionList []model.Article, err error) {
	global.DB.Table("mini_app_article").Find(&CompetitionList)
	return CompetitionList, nil
}
