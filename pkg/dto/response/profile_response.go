package response

type ProfileResponse struct {
	ID uint `json:"id"`
	UserID uint `json:"user_id"`
	Age int `json:"age"`
	Sex string `json:"sex"`
	Phone string `json:"phone"`
}