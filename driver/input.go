package driver

import "encoding/json"

type AddDriverInput struct {
	Id            json.Number `json:"id" binding:"required,number"`
	DriverName    string      `json:"name" binding:"required"`
	PhoneNumber   string      `json:"made_in"`
	IdCard        string      `json:"id_card"`
	DriverLicense string      `json:"driver_license"`
	Status        string      `json:"status"`
}
