package main

import (
	"SCIProj/core"
	"SCIProj/global"
	"SCIProj/initialize"
	"go.uber.org/zap"
)

func main() {
	global.VP = core.Viper()
	global.LOG = core.Zap()

	global.DB = initialize.Gorm()

	db, _ := global.DB.DB()
	defer func() {
		err := db.Close()
		if err != nil {
			global.LOG.Error("Database close failed", zap.Any(" err:", err))
		}
	}()

	core.RunWindowsServer()
}
