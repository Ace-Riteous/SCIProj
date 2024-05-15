package dao

import (
	"SCIProj/dto"
	"SCIProj/model"
	"errors"
)

var teamDao *TeamDao

type TeamDao struct {
	BaseDao
}

func NewTeamDao() *TeamDao {
	if teamDao == nil {
		teamDao = &TeamDao{
			NewBaseDao(),
		}
	}
	return teamDao
}

func (m TeamDao) GetTeamAll(page int, size int) ([]model.Team, int64, error) {
	var teamList []model.Team
	var nTotal int64
	err := m.Orm.Model(&model.Team{}).
		Count(&nTotal).
		Offset((page - 1) * size).
		Limit(size).
		Find(&teamList).
		Error
	if err != nil {
		return nil, 0, err
	}
	return teamList, nTotal, nil
}

func (m TeamDao) GetTeamNotFull(page int, limit int) ([]model.Team, int64, error) {
	var teamList []model.Team
	var nTotal int64
	err := m.Orm.Model(&model.Team{}).
		Where("is_full = ?", false).
		Count(&nTotal).
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&teamList).
		Error
	if err != nil {
		return nil, 0, err
	}
	return teamList, nTotal, nil

}

func (m TeamDao) NewTeam(dto dto.TeamAddDTO, uid int64) error {
	var team model.Team
	team = dto.Convert(uid, false)
	if err := m.Orm.Model(&model.Team{}).Save(&team).Error; err != nil {
		return err
	}
	return nil
}

func (m TeamDao) JoinTeam(joinDTO dto.TeamJoinDTO, uid int64) error {
	var team, team1 model.Team
	err := m.Orm.Model(&model.Team{}).Where("id = ?", joinDTO.TeamID).First(&team).Error
	if err != nil {
		return err
	}
	team.IsFull, err = m.TeamIsFull(joinDTO.TeamID)
	if err != nil {
		return err
	}
	if team.IsFull {
		return errors.New("team is full")
	}
	{
		team1.Name = team.Name
		team1.TeacherID = team.TeacherID
		team1.CID = team.CID
		team1.StudentID = uid
		team1.IsFull = team.IsFull
	}
	err = m.Orm.Model(&model.Team{}).Create(&team1).Error
	if err != nil {
		return err
	}
	return nil
}

func (m TeamDao) TeamIsFull(id int) (bool, error) {
	var post model.Competition
	var team model.Team
	var nTotal int64
	err := m.Orm.Model(&model.Team{}).Where("id = ?", id).First(&team).Error
	if err != nil {
		return false, err
	}
	err = m.Orm.Model(&model.Competition{}).Where("id = ?", team.CID).First(&post).Error
	if err != nil {
		return false, err
	}
	err = m.Orm.Model(&model.Team{}).
		Where("c_id = ? and teacher_id = ? and team_name =?", team.CID, team.TeacherID, team.Name).
		Count(&nTotal).
		Error
	if err != nil {
		return false, err
	}
	if int(nTotal) >= post.Member {
		err = m.Orm.Model(&model.Team{}).Where("id = ?", id).Update("is_full", true).Error
		return true, nil
	}
	return false, nil
}
