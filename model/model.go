package model

// Basic model definitions

import (
	"time"

	"gorm.io/gorm"
)

var DurationOneDay, _ = time.ParseDuration("24h")

type BlackListedDays uint8

// `Time` specifies time in seconds since 00:00 on the given Weekday
// `Weekday` specifies day of week starting from sunday which is 0
type WeeklyAtTime struct {
	Weekday uint8 `gorm:"not null"`
	Time    int32 `gorm:"not null"`
}

type model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	model
	UserID    string `gorm:"primaryKey;not null;unique"`
	DiscordID string
	Name      string `gorm:"not null"`
}

type Timeslot struct {
	model
	TimeslotID string `gorm:"primaryKey;not null;unique"`
	WeeklyAtTime
	UserID string
	User   User
}

type Task struct {
	model
	TaskID      string `gorm:"primaryKey;not null;unique"`
	Description string
	Due         time.Time
	UserID      string `gorm:"index"`
	User        User
	TimeslotID  string
	Timeslot    Timeslot
	BlackListedDays
}
