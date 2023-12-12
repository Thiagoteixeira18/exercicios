package exercicios

import (
	"fmt"
	"time"
)

// Schedule function parses a textual representation of an appointment date into the corresponding time.Time format.
func Schedule(appt string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, appt)
	return t
}

// HasPassed function checks if the appointment was in the past.
func HasPassed(appt string) bool {
	layout := "January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, appt)
	return t.Before(time.Now())
}

// IsAfternoonAppointment function checks if the appointment is in the afternoon (>= 12:00 and < 18:00).
func IsAfternoonAppointment(appt string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, appt)
	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description function returns a description of the appointment date and time.
func Description(appt string) string {
	t := Schedule(appt)
	weekday := t.Weekday().String()
	month := t.Month().String()
	day := t.Day()
	year := t.Year()
	hour := t.Hour()
	minute := t.Minute()

	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %02d:%02d.", weekday, month, day, year, hour, minute)
}

// AnniversaryDate function returns the anniversary date of the salon's opening for the current year in UTC.
func AnniversaryDate() time.Time {
	year := time.Now().Year()
	return time.Date(year, time.September, 15, 0, 0, 0, 0, time.UTC)
}
