package task

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Uid       uint   `gorm:"column:uid;index;not null"`
	Title     string `gorm:"column:title;type:varchar(255);not null"`
	Status    int    `gorm:"column:status;default:0"`
	Content   string `gorm:"column:content;type:longtext"`
	StartTime int64  `gorm:"column:start_time;default:0"`
	EndTime   int64  `gorm:"column:end_time;default:0"`
}
