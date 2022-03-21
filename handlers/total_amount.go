package costs

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/tusharhow/bou-cost/db"
	co "github.com/tusharhow/bou-cost/models"
	"go.mongodb.org/mongo-driver/bson"
)

func TotalAmount(c *gin.Context) {
	collection := db.MGI.Db.Collection("cost")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var totalAmount []co.Cost
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "status": "error"})
		return
	}
	err = cursor.All(ctx, &totalAmount)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "status": "error"})
		return
	}
	var total float64
	for _, amount := range totalAmount {
		total += amount.Cost
	}
	c.JSON(200, gin.H{"status": "success", "data": total})
}
