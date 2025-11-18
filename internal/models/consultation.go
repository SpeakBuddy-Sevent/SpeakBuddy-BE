package models

import "time"

type Consultation struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	UserID          uint      `json:"user_id"`
	TherapistUserID uint      `json:"therapist_user_id"`

	ChildName string `json:"child_name"`
	ChildAge  uint   `json:"child_age"`
	ChildSex  string `json:"child_sex"`

	Date     time.Time `json:"date"`
	TimeSlot string    `json:"time_slot"`
	IsPaid   bool      `json:"is_paid" gorm:"default:false"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User     User `json:"user" gorm:"foreignKey:UserID"`
	Therapist User `json:"therapist" gorm:"foreignKey:TherapistUserID"`
}
