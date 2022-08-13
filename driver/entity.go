package driver

import (
	"time"
	"gorm.io/gorm"
)

type Driver struct {
	gorm.Model
	Id            int
	DriverName    string
	PhoneNumber   string
	IdCard        string
	DriverLicense string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
