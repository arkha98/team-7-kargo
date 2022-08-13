package shipment

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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

func (c *ShipmentController) Update(context *gin.Context) {

	// shipmentId := context.Param("id")

	// shipment :=
	// jsonData, err := ioutil.ReadAll(context.Request.Body)
	// if err != nil {
	// 	// Handle error
	// }

	// var params Shipment
	// err = json.Unmarshal(jsonData, &params)

	// result := c.Database.Create(&params)

	// fmt.Println(params)
	// if result.Error != nil {
	// 	context.JSON(500, "error")
	// }
	context.JSON(200, params)
}
