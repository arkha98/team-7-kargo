package truck

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Truck, error)
	AddTruck(truck Truck) (Truck, error)
	UpdateTruck(truck Truck) (Truck, error)
	DeleteTruck(truck Truck) (Truck, error)
	GetTruckByID(Truck Truck) (Truck, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
