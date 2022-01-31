package model

// Create new instances of models

import "github.com/google/uuid"

// Create new BlackListedDays object
func NewBlackListedDays() BlackListedDays {
	return 0
}

// Create new model object
func newModel() model {
	return model{Id: uuid.NewString()}
}

// Create new User object
func NewUser() User {
	return User{
		model: newModel(),
	}
}

// Create new WeeklyAtTime object
func NewWeeklyAtTime(Weekday uint8, Time int64) WeeklyAtTime {
	return WeeklyAtTime{Weekday, Time}
}

// Create new Timeslot object
func NewTimeslot(time WeeklyAtTime, user User) Timeslot {
	return Timeslot{
		model:        newModel(),
		WeeklyAtTime: time,
		User:         user,
	}
}
