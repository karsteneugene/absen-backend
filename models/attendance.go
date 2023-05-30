package models

import "time"

type Attendance struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	DateCreated time.Time `json:"date_created"`
	ClockIn     time.Time `json:"clock_in"`
	ClockOut    time.Time `json:"clock_out"`
	Location    string    `json:"location"`
	Image       Image     `json:"-" gorm:"foreignKey:ID"`
	User        User      `json:"-" gorm:"foreignKey:ID"`
}
