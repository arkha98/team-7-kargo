package shipment

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(router *gin.RouterGroup, db *gorm.DB) {
	shipment := router.Group("/shipment")

	controller := ShipmentController{Database: db}

	shipment.GET("/", controller.ShipmentAll)
}
