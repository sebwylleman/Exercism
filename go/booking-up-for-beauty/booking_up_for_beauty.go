package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	const format = "1/2/2006 15:04:05"
	t, _ := time.Parse(format, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	const format = "January 2, 2006 15:04:05"
	t, _ := time.Parse(format, date)
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	const format = "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(format, date)
	return t.Hour() >= 12 && t.Hour() <= 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	const format = "Monday, January 2, 2006, at 15:04"
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return fmt.Sprintf("You have an appointment on %s.", t.Format(format))

}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
