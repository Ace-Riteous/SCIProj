package dao

import (
	"SCIProj/dto"
	"SCIProj/model"
)

var postDao *PostDao

type PostDao struct {
	BaseDao
}

func NewPostDao() *PostDao {
	if postDao == nil {
		postDao = &PostDao{
			NewBaseDao(),
		}
	}
	return postDao
}

func (m PostDao) GetCompetitionAll(page int, limit int) ([]model.Competition, int64, error) {
	var competitionList []model.Competition
	var nTotal int64
	err := m.Orm.Model(&model.Competition{}).
		Count(&nTotal).
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&competitionList).
		Error
	if err != nil {
		return nil, 0, err
	}
	return competitionList, nTotal, nil
}

func (m PostDao) CheckIsTeacher(uid int64) bool {
	var teacher model.Teacher
	err := m.Orm.Model(&model.Teacher{}).
		Where("seven_id = ?", uid).
		First(&teacher).
		Error
	if err != nil {
		return false
	}
	return true
}

func (m PostDao) AddCompetition(dto dto.CompetitionAddDTO) error {
	var competition model.Competition
	competition = dto.Convert()
	if err := m.Orm.Model(&model.Competition{}).Save(&competition).Error; err != nil {
		return err
	}
	return nil

}
