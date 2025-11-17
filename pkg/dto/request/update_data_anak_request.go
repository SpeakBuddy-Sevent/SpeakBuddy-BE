package request

type UpdateDataAnakRequest struct {
	ChildName string `json:"child_name"`
	ChildAge int `json:"child_age"`
	ChildSex string `json:"child_sex"`
}