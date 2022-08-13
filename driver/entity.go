package driver

import (
	"gorm.io/gorm"
	"time"
)

type Driver struct {
	gorm.Model    `json:"-"`
	ID            int    `json:"id"`
	DriverName    string `json:"name" binding:"required"`
	PhoneNumber   string `json:"phone_number"`
	IdCard        string `json:"id_card"`
	DriverLicense string `json:"driver_license"`
	Status        string `json:"status"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
