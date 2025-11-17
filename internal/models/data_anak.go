package models

type DataAnak struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `gorm:"uniqueIndex" json:"user_id"`
	ChildName  string `json:"child_name"`
	ChildAge  int    `json:"child_age"`
	ChildSex string `json:"child_sex"`
}