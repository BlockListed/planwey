package model

// Create new instances of models
import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

// Create new BlackListedDays object
func NewBlackListedDays() BlackListedDays {
	return 0
}

// Create new User object
func NewUser() User {
	id, _ := gonanoid.New()
	return User{
		UserID: id,
	}
}

// Create new WeeklyAtTime object
func NewWeeklyAtTime(Weekday uint8, Time int32) WeeklyAtTime {
	return WeeklyAtTime{
		Weekday,
		Time,
	}
}

// Create new Timeslot object
func NewTimeslot(time WeeklyAtTime, userId string) Timeslot {
	id, _ := gonanoid.New()
	return Timeslot{
		TimeslotID:   id,
		WeeklyAtTime: time,
		UserID:       userId,
	}
}

// Create new Task object
func NewTask(description string, due time.Time, userId string,
	timeslotId string, blacklisted BlackListedDays) Task {
	id, _ := gonanoid.New()
	return Task{
		TaskID:          id,
		Description:     description,
		Due:             due,
		UserID:          userId,
		TimeslotID:      timeslotId,
		BlackListedDays: blacklisted,
	}
}
