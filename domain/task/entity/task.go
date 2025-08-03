package entity

import (
	"time"
)

type Task struct {
	Id        uint      `json:"id"`
	Uid       uint      `json:"uid"`
	UserName  string    `json:"user_name"`
	Title     string    `json:"title"`
	Status    int       `json:"status"`
	Content   string    `json:"content"`
	StartTime int64     `json:"start_time"`
	EndTime   int64     `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
