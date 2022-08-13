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
