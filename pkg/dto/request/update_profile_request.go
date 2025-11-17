package request

type UpdateProfileRequest struct {
	Age int `json:"age"`
	Sex string `json:"sex"`
	Phone string `json:"phone"`
}