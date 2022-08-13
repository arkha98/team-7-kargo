package truck

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(router *gin.RouterGroup, db *gorm.DB) {
	truck := router.Group("/truck")

	// truckRepo := NewRepository(db)

	// truck.GET("/query-truck", truckAll)
	// truck.GET("/query-truck/:id", truckId)
	// truck.GET("/query-harga", truckHarga)
	// truck.POST("/add-truck", addtruck)
	// truck.DELETE("/delete-truck", deletetruck)
}
