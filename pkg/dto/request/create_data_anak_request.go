package request

type CreateDataAnakRequest struct {
	ChildName string `json:"child_name"`
	ChildAge int `json:"child_age"`
	ChildSex string `json:"child_sex"`
}