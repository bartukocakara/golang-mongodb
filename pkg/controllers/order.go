package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"restaurant-management/pkg/database"
	"restaurant-management/pkg/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		
		result, err := orderCollection.Find(context.TODO(), bson.M{})
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occured while listing order items"})
		}
		var allOrders []bson.M
		if err := result.All(ctx, &allOrders); err != nil {
			log.Fatal(err)
		}
		
		c.JSON(http.StatusOK, gin.H{"items": allOrders})
	}
}


// ShowOrder godoc
// @Summary Show an Order
// @Description get string by ID
// @Tags accounts
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Router /order/{order_id} [get]
func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		orderId := c.Query("order_id")
		var order models.Order

		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)

		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error" : "error occured while fetching the order item"})
		}
		c.JSON(http.StatusOK, order)
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var table models.Table
		var order models.Order
		
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		order.Order_Date, _ = time.Parse(time.RFC3339, order.Order_Date.Format(time.RFC3339))
		validationErr := validate.Struct(order)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		if order.Table_id != nil {
			err := tableCollection.FindOne(ctx, bson.M{"table_id": order.Table_id}).Decode(&table)
			defer cancel()
			if err != nil {
				msg := fmt.Sprintf("message: Table was not found")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
		}
		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		order.ID = primitive.NewObjectID()
		order.Order_id = order.ID.Hex()

		result, err :=orderCollection.InsertOne(ctx, order)
		if err != nil {
			msg := fmt.Sprintf("order item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, result)
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var table models.Table
		var order models.Order

		var updateObj primitive.D

		orderId := c.Param("order_id")

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if order.Table_id != nil {
			err := menuCollection.FindOne(ctx, bson.M{"table_id": table.Table_id}).Decode(&table)
			defer cancel()
			if err != nil {
				msg := fmt.Sprintf("Error decoding order")
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}

			updateObj = append(updateObj, bson.E{"menu", order.Table_id})
		}

		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{"updated_at", order.Updated_at})

		upsert := true

		filter := bson.M{"order_id": orderId}
		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		result, err := orderCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{"$set", updateObj},
			},
			&opt,
		)

		if err != nil {
			msg := fmt.Sprintf("order item updated failed")
			c.JSON(http.StatusInternalServerError, msg)
			return
		}

		defer cancel()
		c.JSON(http.StatusOK, result)

	}
}

func OrderItemOrderCreator(order models.Order) string {
	order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.ID = primitive.NewObjectID()
	order.Order_id = order.ID.Hex()
	
	orderCollection.InsertOne(ctx, order)
	defer cancel()

	return order.Order_id
}

func DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}