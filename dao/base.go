package dao

import (
	"SCIProj/global"
	"gorm.io/gorm"
)

type BaseDao struct {
	//操作数据库
	Orm *gorm.DB
}

func NewBaseDao() BaseDao {
	return BaseDao{
		Orm: global.DB,
	}
}
