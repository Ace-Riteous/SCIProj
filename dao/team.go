package dao

import (
	"SCIProj/global"
	"SCIProj/model"
	"errors"
	"gorm.io/gorm"
	"strings"
)

func TeamIsFull(num int, teamid string) (bool, error) {
	var team model.Team
	err := global.DB.Model(&model.Team{}).Where("teamid = ?", teamid).First(&team).Error
	if err != nil {
		return false, err
	}
	studentList := strings.Split(team.StudentIds, ",")
	if len(studentList) == num {
		return true, nil
	}
	return false, err
}

func NewTeam(team *model.Team) error {
	err := global.DB.Model(&model.Team{}).Create(&team).Error
	if err != nil {
		return err
	}
	return nil
}

func CheckTeamIdExist(s string) (bool, error) {
	var team model.Team
	err := global.DB.Model(&model.Team{}).Where("teamid = ?", s).First(&team).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return true, err
	}
	return true, errors.New("TeamId已存在")

}

func FetchTeamList() (teamList []model.Team, err error) {
	err = global.DB.Model(&model.Team{}).Find(&teamList).Error
	if err != nil {
		return nil, err
	}
	if len(teamList) == 0 {
		return nil, errors.New("没有查询到队伍信息")
	}
	return teamList, nil

}

func FetchTeamNotFullList() (teamlist []model.Team, err error) {
	err = global.DB.Model(&model.Team{}).Where("is_full = ?", false).Find(&teamlist).Error
	if err != nil {
		return nil, err
	}
	if len(teamlist) == 0 {
		return nil, errors.New("所有队伍已满员")
	}
	return teamlist, nil
}

func FetchTeamByTeamId(teamid string) (team model.Team, err error) {
	err = global.DB.Model(&model.Team{}).Where("teamid = ?", teamid).First(&team).Error
	if err != nil {
		return team, err
	}
	return team, nil
}

func UpdateTeam(team model.Team) error {
	err := global.DB.Model(&model.Team{}).Where("teamid = ?", team.TeamId).Updates(&team).Error
	if err != nil {
		return err
	}
	return nil

}
