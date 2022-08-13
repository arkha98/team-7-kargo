package shipment

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	driver "kargo-tms/driver"
	truck "kargo-tms/truck"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShipmentController struct {
	Database *gorm.DB
}

func (c *ShipmentController) ShipmentAll(context *gin.Context) {
	shipments := []Shipment{}
	result := c.Database.Preload("Driver").Preload("Truck").Find(&shipments)
	jsonData, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		// Handle error
	}
	fmt.Println(jsonData)
	if result.Error != nil {
		context.JSON(500, "testlagi")
	}
	context.JSON(200, shipments)
}

func (c *ShipmentController) Create(context *gin.Context) {
	jsonData, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		// Handle error
	}

	var params Shipment
	err = json.Unmarshal(jsonData, &params)

	result := c.Database.Create(&params)

	fmt.Println(params)
	if result.Error != nil {
		context.JSON(500, "error")
	}
	context.JSON(200, params)
}

type UpdateData struct {
	ID             *int           `json:"id"`
	Origin         *string        `json:"origin"`
	Destination    *string        `json:"destination"`
	LoadingDate    *time.Time     `json:"loading_date"`
	Status         *string        `json:"status"`
	ShipmentNumber *string        `json:"shipment_number"`
	Truck          *truck.Truck   `json:"truck"`
	TruckID        *int           `json:"truck_id"`
	Driver         *driver.Driver `json:"driver"`
	DriverID       *int           `json:"driver_id"`
}

func (c *ShipmentController) Update(context *gin.Context) {

	jsonData, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		// Handle error
	}

	var params UpdateData
	err = json.Unmarshal(jsonData, &params)
	shipment := Shipment{
		ID: int(*params.ID),
	}

	result := c.Database.First(&shipment)
	if result.Error != nil {
		context.JSON(500, "error")
	}

	updateShipment := false

	if params.ShipmentNumber != nil {
		updateShipment = true
		shipment.ShipmentNumber = *params.ShipmentNumber
	}

	if params.Destination != nil {
		updateShipment = true
		shipment.Destination = *params.Destination
	}

	if params.Origin != nil {
		updateShipment = true
		shipment.Origin = *params.Origin
	}

	if params.TruckID != nil {
		updateShipment = true
		shipment.TruckID = params.TruckID
	}

	if params.DriverID != nil {
		updateShipment = true
		shipment.DriverID = params.DriverID
	}

	if updateShipment {
		result := c.Database.Save(&shipment)
		if result.Error != nil {
			context.JSON(500, "error")
		}
	}

	context.JSON(200, params)
}

func (c *ShipmentController) Delete(context *gin.Context) {

	jsonData, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		// Handle error
	}

	var params UpdateData
	err = json.Unmarshal(jsonData, &params)
	shipment := Shipment{
		ID: int(*params.ID),
	}

	result := c.Database.First(&shipment)
	if result.Error != nil {
		context.JSON(500, "error")
	}

	updateShipment := false

	if params.ShipmentNumber != nil {
		updateShipment = true
		shipment.ShipmentNumber = *params.ShipmentNumber
	}

	if params.Destination != nil {
		updateShipment = true
		shipment.Destination = *params.Destination
	}

	if params.Origin != nil {
		updateShipment = true
		shipment.Origin = *params.Origin
	}

	if params.TruckID != nil {
		updateShipment = true
		shipment.TruckID = params.TruckID
	}

	if params.DriverID != nil {
		updateShipment = true
		shipment.DriverID = params.DriverID
	}

	if updateShipment {
		result := c.Database.Save(&shipment)
		if result.Error != nil {
			context.JSON(500, "error")
		}
	}

	context.JSON(200, params)
}
