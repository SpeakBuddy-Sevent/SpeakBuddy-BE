package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `json:"name"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role" gorm:"default:user"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`

	Profile  Profile  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"profile"`
	DataAnak DataAnak `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"data_anak"`
}
