package driver

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DriverController struct {
	Database *gorm.DB
}

func (c *DriverController) DriverAll(context *gin.Context) {
	var drivers []Driver

	_ = c.Database.Debug().Find(&drivers)

	context.JSON(200, drivers)

}

func (c *DriverController) DriverCreate(driver Driver) (Driver, error) {
	err := c.Database.Debug().Create(&driver).Error
	return driver, err
}

func (c *DriverController) CreateDriver(context *gin.Context) {
	jsonData, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(500, "error")
	}

	var params Driver
	err = json.Unmarshal(jsonData, &params)
	result := c.Database.Create(&params)

	if result.Error != nil {
		context.JSON(500, "error")
	}
	context.JSON(200, params)
}

func (c *DriverController) DeleteDriver(context *gin.Context) {
	var params Driver
	var driver Driver

	jsonData, err := ioutil.ReadAll(context.Request.Body)
	err = json.Unmarshal(jsonData, &params)
	if err != nil {
		context.JSON(500, "error")
	}
	c.Database.Debug().Where("id = ?", params.ID).First(&driver)
	c.Database.Debug().Delete(&driver)
}

type UpdateData struct {
	ID            *int    `json:"id"`
	DriverName    *string `json:"name" binding:"required"`
	PhoneNumber   *string `json:"made_in"`
	IdCard        *string `json:"id_card"`
	DriverLicense *string `json:"driver_license"`
	Status        *string `json:"status"`
}

func (c *DriverController) UpdateDriver(context *gin.Context) {

	jsonData, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		context.JSON(500, "error")
	}

	var params UpdateData
	err = json.Unmarshal(jsonData, &params)
	driver := Driver{
		ID: int(*params.ID),
	}

	result := c.Database.First(&driver)
	if result.Error != nil {
		context.JSON(500, "error")
	}

	updateDriver := false

	if params.DriverName != nil {
		updateDriver = true
		driver.DriverName = *params.DriverName
	}

	if params.PhoneNumber != nil {
		updateDriver = true
		driver.PhoneNumber = *params.PhoneNumber
	}

	if params.IdCard != nil {
		updateDriver = true
		driver.IdCard = *params.IdCard
	}

	if params.DriverLicense != nil {
		updateDriver = true
		driver.DriverLicense = *params.DriverLicense
	}

	if params.Status != nil {
		updateDriver = true
		driver.Status = *params.Status
	}

	if updateDriver {
		result := c.Database.Save(&driver)
		if result.Error != nil {
			context.JSON(500, "error")
		}
	}

	context.JSON(200, params)
}
