package model

import "time"

type Task struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Title      string    `json:"title" gorm:"not null"`
	Created_At time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated-at"`
	User       User      `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	UserID     uint      `json:"user_id" gorm:"not null"`
}

type TaskResponse struct {
	ID         uint      `json:"id"`
	Title      string    `json:"title"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"updated_at"`
}
