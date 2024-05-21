package dao

import (
	"SCIProj/dto"
	"SCIProj/model"
	"errors"
	"strconv"
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
		Offset((page - 1) * size).
		Limit(size).
		Find(&teamList).
		Offset(-1).
		Limit(-1).
		Count(&nTotal).
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
		Offset((page - 1) * limit).
		Limit(limit).
		Find(&teamList).
		Offset(-1).
		Limit(-1).
		Count(&nTotal).
		Error
	if err != nil {
		return nil, 0, err
	}
	return teamList, nTotal, nil

}

func (m TeamDao) NewTeam(dto dto.TeamAddDTO, uid string) error {
	var team model.Team
	var post model.Competition
	err := m.Orm.Model(&model.Competition{}).Where("id = ?", dto.CID).First(&post).Error
	team = dto.Convert(uid, false, post.Member)
	if CheckHasTeamed(team) {
		return errors.New("every student can only create one team in one competition")
	}
	if err = m.Orm.Model(&model.Team{}).Save(&team).Error; err != nil {
		return err
	}
	return nil
}

func CheckHasTeamed(team model.Team) bool {
	var tempTeam model.Team
	err := teamDao.Orm.Model(&model.Team{}).
		Where("c_id = ? and student_id = ?", team.CID, team.StudentID).
		First(&tempTeam).
		Error
	if err != nil {
		return false
	}
	return true
}

func (m TeamDao) JoinTeam(joinDTO dto.TeamJoinDTO, uid string) error {
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
	if CheckInTeam(joinDTO.TeamID, uid) {
		return errors.New("you are already in the team")
	}
	{
		team1.TeamName = team.TeamName
		team1.TeacherID = team.TeacherID
		team1.CID = team.CID
		team1.StudentID = uid
		team1.LeaderID = team.LeaderID
		team1.QQGroup = team.QQGroup
		team1.IsFull = team.IsFull
		team1.Now2all = team.Now2all
	}
	err = m.Orm.Model(&model.Team{}).Create(&team1).Error
	if err != nil {
		return err
	}
	return nil
}

func CheckInTeam(teamID int, uid string) bool {
	var team model.Team
	err := teamDao.Orm.Model(&model.Team{}).
		Where("id = ? and student_id = ?", teamID, uid).
		First(&team).
		Error
	if err != nil {
		return false
	}
	return true
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
		Where("c_id = ? and teacher_id = ? and team_name =?", team.CID, team.TeacherID, team.TeamName).
		Count(&nTotal).
		Error
	if err != nil {
		return false, err
	}
	if int(nTotal) >= post.Member {
		err = m.Orm.Model(&model.Team{}).Where("id = ?", id).Update("is_full", true).Error
		if err != nil {
			return true, err
		}
		n2a := strconv.Itoa(int(nTotal)) + "/" + strconv.Itoa(post.Member)
		err = m.Orm.Model(&model.Team{}).Where("id = ?", id).Update("now2all", n2a).Error
		if err != nil {
			return true, err
		}
		return true, nil
	}
	n2a := strconv.Itoa(int(nTotal)) + "/" + strconv.Itoa(post.Member)
	err = m.Orm.Model(&model.Team{}).Where("id = ?", id).Update("now2all", n2a).Error
	return false, nil
}
