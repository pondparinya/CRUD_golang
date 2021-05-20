package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type StudentEntity struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	StudentID string             `json:"student_id" bson:"student_id,omitempty"`
	Name      string             `json:"name" bson:"name,omitempty"`
	Nickname  string             `json:"nickname" bson:"nickname,omitempty"`
	Age       string             `json:"age" bson:"age,omitempty"`
	Address   *AddressEntity     `json:"address" bson:"address,omitempty"`
}

type AddressEntity struct {
	StreetAddress  string `json:"streetAddress" bson:"streetAddress,omitempty"`
	StreetAddress2 string `json:"streetAddress2" bson:"streetAddress2,omitempty"`
	City           string `json:"city" bson:"city,omitempty"`
	State          string `json:"state" bson:"state,omitempty"`
	Zipcode        string `json:"zipcode" bson:"zipcode,omitempty"`
	Country        string `json:"country" bson:"country,omitempty"`
}

