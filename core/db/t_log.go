package db

import "time"

type TLog struct {
	Id        int       `json:"id" gorm:"autoIncrement; primaryKey"`
	User      string    `json:"user"`
	Level     int8      `json:"level"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
