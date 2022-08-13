package shipment

import (
	driver "kargo-tms/driver"
	truck "kargo-tms/truck"
	"time"

	"gorm.io/gorm"
)

type Shipment struct {
	gorm.Model
	Id             int
	Origin         string
	Destination    string
	LoadingDate    time.Time
	Status         string
	ShipmentNumber string
	Truck          truck.Truck
	TruckID        int
	Driver         driver.Driver
	DriverID       int
}
