package driver

import (
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
