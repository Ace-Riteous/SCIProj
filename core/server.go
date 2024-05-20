package core

import (
	"SCIProj/initialize"
	"SCIProj/router"
	"SCIProj/service"
)

func RunWindowsServer() {
	go service.ServerHealthCheck()
	//initialize.Redis()
	initialize.JWTANDMD()
	//调试时使用：打印当前时间的int64值
	//fmt.Println(time.Now().Unix())
	newData()
	router.InitRouter()
}
