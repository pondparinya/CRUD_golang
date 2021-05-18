package dao

type StudentDAO struct {
	ID        string     `json:"_id"`
	StudentID int        `json:"student_id" validate:"required"`
	Name      string     `json:"name" validate:"required"`
	Nickname  string     `json:"nickname,omitempty"`
	Age       int        `json:"age" validate:"required"`
	Address   AddressDAO `json:"address" validate:"dive"`
}


type AddressDAO struct {
	StreetAddress  string `json:"streetAddress"`
	StreetAddress2 string `json:"streetAddress2"`
	City           string `json:"city"`
	State          string `json:"state"`
	Zipcode        string `json:"zipcode"`
	Country        string `json:"country"`
}
