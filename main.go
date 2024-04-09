package main

import (
	Routers "SCIProj/router"
	"log"
)

func main() {
	r := Routers.SetRouters()
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server start failed")
	}
}
