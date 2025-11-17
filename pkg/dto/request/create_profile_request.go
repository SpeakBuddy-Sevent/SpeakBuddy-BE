package request

type CreateProfileRequest struct {
	Age int `json:"age"`
	Sex string `json:"sex"`
	Phone string `json:"phone"`
}