package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// ToDoList struct
type ToDoList struct {
	primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task               string `json:"task,omitempty"`
	Status             bool   `json:"status,omitempty"`
}
