package service

import (
	"SCIProj/utils"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"math/rand"
	"os"
	"strconv"
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

	timeNow := time.Now().Unix()
	HighCpuUsage := rand.Float64()*100 + 1

	slog, err := os.Create("./service_log/service_log" + strconv.Itoa(int(timeNow)) + ".txt")
	if err != nil {
		panic(err)
	}
	//写入日志
	defer slog.Close()

	_, err = slog.WriteString("[SCIProj ->]service is running..." + time.Now().String() + "\n")
	_, err = slog.WriteString("\t\t\tWarning cpu usage: " + fmt.Sprintf("%v\n", HighCpuUsage))

	percentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		_, _ = slog.WriteString("error: " + err.Error() + "\n")
		panic(err)
	}

	timeTick := time.NewTicker(time.Second)
	for {
		select {
		case <-timeTick.C:
			_, err = slog.WriteString("[SCIProj ->]service running normally..." + time.Now().String() + "\n")
			_, err = slog.WriteString("\t\t\tWarning cpu usage: " + fmt.Sprintf("%v\n", HighCpuUsage))
			if err != nil {
				_, _ = slog.WriteString("server error! err: " + err.Error() + "\n")
				panic(err)
			}
			_, err = slog.WriteString(fmt.Sprintf("\t\t\tcpu usage: %v\n", percentages))
			if err != nil {
				_, _ = slog.WriteString("server error! err: " + err.Error() + "\n")
				panic(err)
			}
			percentages, err = cpu.Percent(time.Second, false)
			if err != nil {
				_, _ = slog.WriteString("error: " + err.Error() + "\n")
				panic(err)
			}
			if percentages[0] > utils.GenerateCpuUsage(timeNow, HighCpuUsage) {
				_, _ = slog.WriteString("cpu usage is too high! " + fmt.Sprintf("cpu usage: %v\n", percentages))
				os.Exit(0)
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
	var length int
	if len(uid) < 6 {
		length = 6
	} else {
		length = len(uid)
	}
	//新建一个user_log.txt
	ulog, err := os.Create("./user_log/user_" + utils.GenerateRandomString(length) + uid + "_log.txt")
	if err != nil {
		panic(err)
	}
	//写入日志
	defer ulog.Close()
	_, err = ulog.WriteString("[SCIProj ->]user server is running..." + time.Now().String() + "\n")
	percentages, err := cpu.Percent(time.Second, false)

	if err != nil {
		_, _ = ulog.WriteString("error: " + err.Error() + "\n")
		panic(err)
	}

	timeTick := time.NewTicker(time.Second)
	for {
		select {
		case <-timeTick.C:
			_, err = ulog.WriteString("[SCIProj ->]user server running normally..." + time.Now().String() + "\n")
			if err != nil {
				_, _ = ulog.WriteString("server error! err: " + err.Error() + "\n")
				panic(err)
			}
			_, err = ulog.WriteString(fmt.Sprintf("\t\t\tcpu usage: %v\n", percentages))
			if err != nil {
				_, _ = ulog.WriteString("server error! err: " + err.Error() + "\n")
				panic(err)
			}
			percentages, err = cpu.Percent(time.Second, false)
			if err != nil {
				_, _ = ulog.WriteString("error: " + err.Error() + "\n")
				panic(err)
			}
		}
	}
}
