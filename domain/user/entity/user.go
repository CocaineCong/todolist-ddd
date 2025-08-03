package entity

import (
	"time"
)

type User struct {
	ID        uint
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
