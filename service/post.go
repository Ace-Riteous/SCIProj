package service

import (
	"SCIProj/dao"
	"SCIProj/model"
)

func FetchCompetitionList() (CompetitionList []model.Article, err error) {
	CompetitionList, err = dao.GetCompetitionAll()
	if err != nil {
		return nil, err
	}
	return CompetitionList, nil
}

func AddCompetition(article model.Article) error {
	return nil
}
