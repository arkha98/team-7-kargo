package shipment

import (
	driver "kargo-tms/driver"
	truck "kargo-tms/truck"
	"time"

	"gorm.io/gorm"
)

type Shipment struct {
	gorm.Model
	ID             int           `json:"id"`
	Origin         string        `json:"origin"`
	Destination    string        `json:"destination"`
	LoadingDate    time.Time     `json:"loading_date"`
	Status         string        `json:"status"`
	ShipmentNumber string        `json:"shipment_number"`
	Truck          truck.Truck   `json:"truck"`
	TruckID        int           `json:"truck_id"`
	Driver         driver.Driver `json:"driver"`
	DriverID       int           `json:"driver_id"`
}
