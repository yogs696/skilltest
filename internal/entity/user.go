package entity

import "time"

type (

	// User entity.
	User struct {
		ID        uint64    `gorm:"primaryKey" json:"id"`
		Username  string    `gorm:"not null;size:25;" json:"username"`
		Email     string    `gorm:"not null;size:50;unique;index" json:"email"`
		Password  string    `json:"password" binding:"required"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
