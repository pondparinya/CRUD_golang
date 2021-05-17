package entity

type StudentEntity struct {
	StudentID int            `json:"student_id" bson:"student_id"`
	Name      string         `json:"name" bson:"name"`
	Nickname  string         `json:"nickname" bson:"nickname"`
	Age       int            `json:"age" bson:"age"`
	Address   *AddressEntity `json:"address" bson:"address"`
}

type AddressEntity struct {
	StreetAddress  string `json:"streetAddress" bson:"streetAddress"`
	StreetAddress2 string `json:"streetAddress2" bson:"streetAddress2"`
	City           string `json:"city" bson:"city"`
	State          string `json:"state" bson:"state"`
	Zipcode        string `json:"zipcode" bson:"zipcode"`
	Country        string `json:"country" bson:"country"`
}
