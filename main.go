package main

import (
	"github.com/gin-gonic/gin"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=172.20.0.3 user=kargo-tms password=qwerpoiu dbname=tms port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&driver.Driver{})
	db.AutoMigrate(&shipment.Shipment{})
	db.AutoMigrate(&truck.Truck{})

	router := gin.Default()

	team7 := gin.Group("/")

	driver.Routes(team7, db)
	shipment.Routes(team7, db)
	truck.Routes(team7, db)

	router.Run(":8080")
}
