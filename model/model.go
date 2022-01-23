package model

// Basic model definitions

import (
	"time"

	"gorm.io/gorm"
)

// `Time` specifies time since 00:00 on the given Weekday
type weeklyAtTime struct {
	Weekday uint8  `gorm:"not null"`
	Time    uint32 `gorm:"not null"`
}

type model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	model
	Id        string `gorm:"primaryKey;not null"`
	DiscordId uint64
	Name      string `gorm:"not null"`
}

type Timeslot struct {
	model
	weeklyAtTime
	Id   string `gorm:"primaryKey;not null"`
	User User   `gorm:"not null"`
}

type Task struct {
	model
	Id          string `gorm:"primaryKey;not null"`
	Description string
	Due         time.Time
	User        User     `gorm:"not null"`
	Timeslot    Timeslot `gorm:"not null"`
}
