package service

import (
	"SCIProj/global"
	"SCIProj/utils"
)

func HealthCheck() {
	ok, err := utils.PathExists("./service_log")
	if err != nil {
		panic(err)
	}
	if !ok {
		err := utils.CreateDir("./service_log")
		if err != nil {
			panic(err)
		}
	}
	//写入日志
	global.LOG.Info("service start")

}
