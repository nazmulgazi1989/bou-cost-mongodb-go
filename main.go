package main

import (
	"log"

	"github.com/gin-gonic/gin"

	db "github.com/tusharhow/bou-cost/db"
	co "github.com/tusharhow/bou-cost/handlers"
)





func main() {
	r := gin.Default()
	db.Connect()

	r.POST("/addCost", co.AddCost)
	r.GET("/totalCost", co.TotalCost)
	r.GET("/totalAmount", co.TotalAmount)
	r.DELETE("/deleteCost/:id", co.DeleteCost)
	r.PUT("/updateCost", co.UpdateCost)
	

	log.Fatal(r.Run(":8080"))

}




