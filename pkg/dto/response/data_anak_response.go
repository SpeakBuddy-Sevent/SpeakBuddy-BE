package response

type DataAnakResponse struct {
	ID uint `json:"id"`
	UserID uint `json:"user_id"`
	ChildName string `json:"child_name"`
	ChildAge int `json:"child_age"`
	ChildSex string `json:"child_sex"`
}