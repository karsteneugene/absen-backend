package models

import "time"

type Log struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DateCreated time.Time `json:"date_created"`
	User        User      `json:"-" gorm:"foreignKey"`
}
