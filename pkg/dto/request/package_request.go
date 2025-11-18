package request

import "time"

type CreateConsultationRequest struct {
	ChildName string    `json:"child_name"`
	ChildAge  uint      `json:"child_age"`
	ChildSex  string    `json:"child_sex"`
	Date      time.Time `json:"date"`
	TimeSlot  string    `json:"time_slot"`
	IsPaid    bool      `json:"is_paid"`
}
