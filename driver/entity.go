package driver

import (
	"gorm.io/gorm"
	"time"
)

type Driver struct {
	gorm.Model
	ID            int    `json:"id"`
	DriverName    string `json:"name" binding:"required"`
	PhoneNumber   string `json:"made_in"`
	IdCard        string `json:"id_card"`
	DriverLicense string `json:"driver_license"`
	Status        string `json:"status"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
