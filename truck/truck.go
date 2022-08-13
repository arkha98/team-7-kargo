package truck

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var truckRepo repository

func Routes(router *gin.RouterGroup, db *gorm.DB) {
	truck := router.Group("/truck")

	truckRepo.db = db

	truck.GET("/", truckAll)
	truck.POST("/create", truckAdd)
	truck.PUT("/update", truckUpdate)
	truck.DELETE("/delete", truckDelete)
}

func truckAll(c *gin.Context) {
	result, _ := truckRepo.FindAll()
	c.JSON(http.StatusOK, result)
}

func truckAdd(context *gin.Context) {
	jsonData, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		// Handle error
	}

	var params Truck
	err = json.Unmarshal(jsonData, &params)

	fmt.Println(params)
	truckRepo.AddTruck(params)
	context.JSON(200, params)
}

func truckUpdate(context *gin.Context) {
	jsonData, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		// Handle error
	}

	var params Truck
	err = json.Unmarshal(jsonData, &params)

	fmt.Println(params)
	truckRepo.UpdateTruck(params)
	context.JSON(200, params)
}

func truckDelete(context *gin.Context) {
	jsonData, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		// Handle error
	}

	var params Truck
	err = json.Unmarshal(jsonData, &params)

	fmt.Println(params)
	truckRepo.DeleteTruck(params)
	context.JSON(200, params)
}
