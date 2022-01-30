package model

// Basic model definitions

import (
	"time"

	"gorm.io/gorm"
)

var DurationOneDay, _ = time.ParseDuration("24h")

// `Time` specifies time in seconds since 00:00 on the given Weekday
// `Weekday` specifies day of week starting from sunday which is 0
type WeeklyAtTime struct {
	Weekday uint8 `gorm:"not null"`
	Time    int64 `gorm:"not null"`
}

type model struct {
	Id        string `gorm:"primaryKey;not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	model
	DiscordId uint64
	Name      string `gorm:"not null"`
}

type Timeslot struct {
	model
	WeeklyAtTime
	User User `gorm:"not null"`
}

type Task struct {
	model
	Description string
	Due         time.Time
	User        User     `gorm:"not null"`
	Timeslot    Timeslot `gorm:"not null"`
}
