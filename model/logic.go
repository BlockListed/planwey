package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/now"
	"github.com/snabb/isoweek"
)

// Logic

// Add day to blacklisted days
func (b *BlackListedDays) AddDay(day uint8) {
	*b |= (1 << day)
}

// Check day from bitfield
func (b *BlackListedDays) CheckDay(day uint8) bool {
	return ((*b & (1 << day)) >> day) == 1
}

// Return begginning of week for time.Time
func beginningOfWeek(t time.Time) time.Time {
	return t.AddDate(0, 0, -int(t.Weekday())).Truncate(DurationOneDay)
}

// Duration from since the begginging of the day.
func (w WeeklyAtTime) Duration() time.Duration {
	t, _ := time.ParseDuration(fmt.Sprint(w.Time) + "s")
	return t
}

// Returns the WeeklyAtTime time in the week when the command is executed
func (w WeeklyAtTime) ThisWeek() time.Time {
	return now.BeginningOfWeek().AddDate(0, 0, int(w.Weekday)).Add(w.Duration())
}

// Returns time of event for a given ISO week / year combo.
func (w WeeklyAtTime) WeekYear(year int, week int) time.Time {
	return isoweek.StartTime(year, week, time.UTC).AddDate(0, 0, int(w.Weekday)).Add(w.Duration())
}
