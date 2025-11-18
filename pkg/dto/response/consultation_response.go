package response

type ConsultationResponse struct {
	ID              uint   `json:"id"`
	UserID          uint   `json:"user_id"`
	TherapistUserID uint   `json:"therapist_user_id"`

	ChildName string `json:"child_name"`
	ChildAge  uint   `json:"child_age"`
	ChildSex  string `json:"child_sex"`

	Date     string `json:"date"`
	TimeSlot string `json:"time_slot"`
	IsPaid   bool   `json:"is_paid"`
}
