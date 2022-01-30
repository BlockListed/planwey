package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/now"
	"github.com/snabb/isoweek"
)

// Logic

// Own BeginningOfWeek implementation for time.Time
func BeginningOfWeek(t time.Time) time.Time {
	return t.AddDate(0, 0, -int(t.Weekday()))
}

// Duration from secs
func (w WeeklyAtTime) Duration() time.Duration {
	t, _ := time.ParseDuration(fmt.Sprint(w.Time) + "s")
	return t
}

// Returns the WeeklyAtTime time in the week when the command is executed
func (w WeeklyAtTime) ThisWeek() time.Time {
	return now.BeginningOfWeek().AddDate(0, 0, int(w.Weekday)).Add(w.Duration())
}

func (w WeeklyAtTime) WeekYear(year int, week int) time.Time {
	return isoweek.StartTime(year, week, time.UTC).AddDate(0, 0, int(w.Weekday)).Add(w.Duration())
}
