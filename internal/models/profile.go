package models

type Profile struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `gorm:"uniqueIndex" json:"user_id"`
	Age       int    `json:"age"`
	Sex		  string `json:"sex"`
	Phone     string `json:"phone"`
}