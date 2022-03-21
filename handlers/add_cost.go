package costs

import (
	"context"
	"time"
	co "github.com/tusharhow/bou-cost/models"
	db "github.com/tusharhow/bou-cost/db"
	"github.com/gin-gonic/gin"
)

func AddCost(c *gin.Context) {
	var cost co.Cost
	c.BindJSON(&cost)
	collection := db.MGI.Db.Collection("cost")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, cost)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "status": "error"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "message": "cost added successfully"})
}