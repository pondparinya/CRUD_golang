package dao

type StudentDAO struct {
	StudentID int      `json:"student_id"`
	Name      string   `json:"name"`
	Nickname  string   `json:"nickname"`
	Age       int      `json:"age"`
	Address   AddressDAO `json:"address"`
}

type AddressDAO struct {
	StreetAddress  string `json:"streetAddress"`
	StreetAddress2 string `json:"streetAddress2"`
	City           string `json:"city"`
	State          string `json:"state"`
	Zipcode        string `json:"zipcode"`
	Country        string `json:"country"`
}
