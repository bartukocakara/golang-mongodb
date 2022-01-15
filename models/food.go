package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name *string 	`json:"name" validate:"required,min=2,max10"`
	Price *float64	`json:"price" validate:"required"`
	FoodImage *string  `json:"food_image" validate:"required"`
	Created_At time.Time `json:"created_created_at"`
	Updated_At time.Time `json:"updated_at"`
	Food_Id string `json:"food_id"`
	Menu_Id *string `json:"menu_id" validate:"required"`
}