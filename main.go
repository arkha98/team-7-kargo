package main

import (
	driver "kargo-tms/driver"
	shipment "kargo-tms/shipment"
	truck "kargo-tms/truck"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-contrib/cors"
)

func main() {
	dsn := "host=172.20.0.2 user=kargo-tms password=qwerpoiu dbname=kargo-tms port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	// dsn := "host=localhost user=kargo-tms password=qwerpoiu dbname=kargo-tms port=5437 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&driver.Driver{})
	db.AutoMigrate(&shipment.Shipment{})
	db.AutoMigrate(&truck.Truck{})

	router := gin.Default()

	router.Use(cors.Default())

	team7 := router.Group("/")

	driver.Routes(team7, db)
	shipment.Routes(team7, db)
	truck.Routes(team7, db)

	router.Run(":8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
