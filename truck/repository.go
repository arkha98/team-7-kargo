package truck

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Truck, error)
	AddTruck(truck Truck) (Truck, error)
	UpdateTruck(truck Truck) (Truck, error)
	DeleteTruck(truck Truck) (Truck, error)
	GetTruckByID(Truck Truck) (Truck, error)
	FindById(id int) (Truck, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Truck, error) {
	var trucks []Truck

	err := r.db.Debug().Find(&trucks).Error

	return trucks, err
}

func (r *repository) AddTruck(truck Truck) (Truck, error) {
	err := r.db.Debug().Create(&truck).Error

	return truck, err
}

func (r *repository) UpdateTruck(truck Truck) (Truck, error) {
	var getTruck Truck
	err := r.db.Debug().Where("id = ?", truck.Id).Take(&getTruck).Error
	getTruck = truck
	err = r.db.Debug().Save(&getTruck).Error

	return getTruck, err
}

func (r *repository) DeleteTruck(truck Truck) (Truck, error) {
	var getTruck Truck
	err := r.db.Debug().Where("id = ?", truck.Id).Take(&getTruck).Error
	err = r.db.Debug().Delete(&getTruck).Error

	return getTruck, err
}

func (r *repository) FindById(id int) (Truck, error) {
	var getTruck Truck
	err := r.db.Debug().Where("id = ?", id).Take(&getTruck).Error

	return getTruck, err
}
