package dao

type CreateStudentDAO struct {
	StudentID int        `json:"student_id" validate:"required"`
	Name      string     `json:"name" validate:"required"`
	Nickname  string     `json:"nickname,omitempty"`
	Age       int        `json:"age" validate:"required"`
	Address   AddressDAO `json:"address" validate:"dive"`
}
type CreateStudentResDAO struct {
	ID string `json:"_id"`
}
type GetStudentResDAO struct {
	ID        string     `json:"_id"`
	StudentID int        `json:"student_id,omitempty" validate:"required"`
	Name      string     `json:"name,omitempty" validate:"required"`
	Nickname  string     `json:"nickname,omitempty"`
	Age       int        `json:"age,omitempty" validate:"required"`
	Address   AddressDAO `json:"address,omitempty" validate:"dive"`
}

type AddressDAO struct {
	StreetAddress  string `json:"streetAddress,omitempty"`
	StreetAddress2 string `json:"streetAddress2,omitempty"`
	City           string `json:"city,omitempty"`
	State          string `json:"state,omitempty"`
	Zipcode        string `json:"zipcode,omitempty"`
	Country        string `json:"country,omitempty"`
}
