package entity

import "time"

type Schedule struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CinemaID  uint      `json:"cinema_id"`
	MovieID   uint      `json:"movie_id"`
	ShowDate  time.Time `json:"show_date"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
