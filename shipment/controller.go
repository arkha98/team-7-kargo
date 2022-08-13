package shipment

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShipmentController struct {
	Database *gorm.DB
}

func (c *ShipmentController) ShipmentAll(context *gin.Context) {
	// votePack, err := c.Database.
	// if err != nil {
	// 	context.String(404, "Votepack Not Found")
	// }
	context.JSON(200, "votePack")
}
