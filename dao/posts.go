package dao

import (
	"SCIProj/global"
	"SCIProj/model"
	"errors"
)

func GetCompetitionAll() (CompetitionList []*model.Article, err error) {
	global.DB.Table("articles").Find(&CompetitionList)
	if len(CompetitionList) == 0 {
		return nil, errors.New("没有比赛信息！")
	}
	return CompetitionList, nil
}

func AddCompetition(article *model.Article) error {
	global.DB.Table("articles").Create(article)
	return nil
}
