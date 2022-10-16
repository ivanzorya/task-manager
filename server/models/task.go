package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID         	primitive.ObjectID 	`bson:"_id"`
	Subject    	*string            	`json:"subject"`
	Done	   	*bool           	`json:"done"`
}
