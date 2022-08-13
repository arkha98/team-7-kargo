package driver

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(router *gin.RouterGroup, db *gorm.DB) {
	driver := router.Group("/driver")
	controller := DriverController{Database: db}

	driver.GET("/", controller.DriverAll)
	driver.POST("/create", controller.CreateDriver)
	driver.POST("/delete", controller.DeleteDriver)
}
