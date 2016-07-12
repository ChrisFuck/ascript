package main

import (
	"github.com/ZachBergh/common/mysql"
)

func init() {
	mysql.InitMysqlConfig(
		"127.0.0.1:3306",
		"127.0.0.1:3306",
		"root",
		"",
		"aeris",
	)
}

func main(){
    deviceList := initDevice()

    n := len(deviceList)

    for i := 0; i < n; i ++ {
		deviceList[i].Icon = "icon-dog"
        UpdateIcon(deviceList[i])
	}
}
