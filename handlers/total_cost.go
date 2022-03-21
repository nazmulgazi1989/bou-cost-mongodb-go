package costs

import (
	"context"
	"time"
	co "github.com/tusharhow/bou-cost/models"
	db "github.com/tusharhow/bou-cost/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func TotalCost(c *gin.Context) {
	collection := db.MGI.Db.Collection("cost")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var totalCost []co.Cost
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "status": "error"})
		return
	}
	err = cursor.All(ctx, &totalCost)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "status": "error"})
		return
	}
	c.JSON(200, gin.H{"status": "success", "data": totalCost})
}