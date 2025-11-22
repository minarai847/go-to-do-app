package model

import "time"

type User struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	Email      string    `json:"email" gorm:"unique"`
	Password   string    `json:"password"`
	Created_At time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated-at"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}
