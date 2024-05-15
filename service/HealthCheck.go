package service

import (
	"SCIProj/utils"
	"os"
	"time"
)

func ServerHealthCheck() {
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
	//新建一个service_log.txt
	slog, err := os.Create("./service_log/service_log.txt")
	if err != nil {
		panic(err)
	}
	//写入日志
	defer slog.Close()
	_, err = slog.WriteString("service is running..." + time.Now().String() + "\n")
	if err != nil {
		panic(err)
	}
	timeTick := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-timeTick.C:
			_, err = slog.WriteString("service running normally..." + time.Now().String() + "\n")
			if err != nil {
				_, _ = slog.WriteString("server error! err: " + err.Error() + "\n")
				panic(err)
			}
		}
	}
}

func UserHealthCheck(uid string) {
	ok, err := utils.PathExists("./user_log")
	if err != nil {
		panic(err)
	}
	if !ok {
		err := utils.CreateDir("./user_log")
		if err != nil {
			panic(err)
		}
	}
	//新建一个user_log.txt
	ulog, err := os.Create("./user_log/user_" + uid + "_log.txt")
	if err != nil {
		panic(err)
	}
	//写入日志
	defer ulog.Close()
	_, err = ulog.WriteString("user server is running..." + time.Now().String() + "\n")
	if err != nil {
		panic(err)
	}
	timeTick := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-timeTick.C:
			_, err = ulog.WriteString("user server running normally..." + time.Now().String() + "\n")
			if err != nil {
				_, _ = ulog.WriteString("server error! err: " + err.Error() + "\n")
				panic(err)
			}
		}
	}
}
