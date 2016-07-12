package main

import (
	"github.com/ZachBergh/common/mysql"
)

const (
	DeviceTable      = "devices_device"
)

type Device struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	Core_Id       string  `json:"core_id"`
	Lat           float64 `json:"lat"`
	Lng           float64 `json:"lng"`
	Register_Date string  `json:"register_date"`
	Last_Updated  string  `json:"last_updated"`
	Group_Id      int     `json:"group_id"`
	User_Id       int     `json:"user_id"`
	Station_Id    int     `json:"station_id"`
	Claim_Code    string  `json:"claim_code"`
	Icon          string  `json:"icon"`
	Firmware      int     `json:"firmware"`
}

type DeviceClient struct {
    Devices []Device
}

func initDevice() []Device{
    d := &DeviceClient{}
    d.Sync()
    return d.Devices
}

func (d *DeviceClient) Data(table string, response interface{}) error {
	mclient := mysql.MysqlQuery{
		Table: table,
	}
	return mclient.FindAll(response, false)
}

func (d *DeviceClient) Sync(){
    var devices []Device
	d.Data(DeviceTable, &devices)
	d.Devices = devices
}

func UpdateIcon(response interface{}){
    mclient := mysql.MysqlQuery{
		Table: DeviceTable,
	}
    mclient.Upsert(response,false);
}
