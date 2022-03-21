package costs

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/tusharhow/bou-cost/db"
	co "github.com/tusharhow/bou-cost/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateCost(c *gin.Context) {
	var cost co.Cost
	c.BindJSON(&cost)
	collection := db.MGI.Db.Collection("cost")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := collection.UpdateOne(ctx, bson.D{{}}, bson.D{{
		"$set", bson.D{

			{"cost", cost.Cost},
			{"costType", cost.CostType},
			{"costDate", cost.CostDate},
			{"costDescription", cost.CostDescription},
		},
	}})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "Cost updated", "message": "Cost updated successfully"})
}
