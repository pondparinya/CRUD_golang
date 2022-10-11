package dao

type CreateStudentReq struct {
	StudentID string     `json:"student_id" validate:"required"`
	Name      string     `json:"name" validate:"required"`
	Nickname  string     `json:"nickname,omitempty"`
	Age       string     `json:"age" validate:"required"`
	Address   AddressDAO `json:"address" validate:"dive"`
}
type CreateStudentRes struct {
	ID string `json:"_id"`
}
type GetStudentRes struct {
	ID        string     `json:"_id"`
	StudentID string     `json:"student_id,omitempty" validate:"required"`
	Name      string     `json:"name,omitempty" validate:"required"`
	Nickname  string     `json:"nickname,omitempty"`
	Age       string     `json:"age,omitempty" validate:"required"`
	Address   AddressDAO `json:"address,omitempty" validate:"dive"`
}

type DeleteStudentReq struct {
	StudentID string `json:"student_id" validate:"required"`
}

type AddressDAO struct {
	StreetAddress  string `json:"streetAddress,omitempty"`
	StreetAddress2 string `json:"streetAddress2,omitempty"`
	City           string `json:"city,omitempty"`
	State          string `json:"state,omitempty"`
	Zipcode        string `json:"zipcode,omitempty"`
	Country        string `json:"country,omitempty"`
}
