package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type StudentEntity struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	StudentID int                `json:"student_id" bson:"student_id"`
	Name      string             `json:"name" bson:"name,omitempty"`
	Nickname  string             `json:"nickname" bson:"nickname"`
	Age       int                `json:"age" bson:"age"`
	Address   *AddressEntity     `json:"address" bson:"address"`
}

type AddressEntity struct {
	StreetAddress  string `json:"streetAddress" bson:"streetAddress"`
	StreetAddress2 string `json:"streetAddress2" bson:"streetAddress2"`
	City           string `json:"city" bson:"city"`
	State          string `json:"state" bson:"state"`
	Zipcode        string `json:"zipcode" bson:"zipcode"`
	Country        string `json:"country" bson:"country"`
}
