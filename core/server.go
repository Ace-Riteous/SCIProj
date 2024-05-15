package core

import (
	"SCIProj/initialize"
	"SCIProj/router"
	"SCIProj/service"
)

func RunWindowsServer() {
	go service.ServerHealthCheck()
	initialize.Redis()
	initialize.JWTANDMD()
	router.InitRouter()
}
