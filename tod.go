package tod

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrInvalidTime   = errors.New("invalid time")
	ErrInvalidString = errors.New("invalid time string")
)

// Time represents a time (HH::MM) in a day
// always in 24H format
type Time struct {
	hours   uint
	minutes uint
}

// NewTime returns a new Time at 00:00
func NewTime() Time {
	return Time{
		hours:   0,
		minutes: 0,
	}
}

// SetTime sets the time to the appropriate hours and minutes
// Time has to be between 0:00 and 23:59
func (t *Time) SetTime(h, m uint) error {
	if h < 0 || h > 23 {
		return ErrInvalidTime
	}
	if m < 0 || m > 59 {
		return ErrInvalidTime
	}
	t.hours = h
	t.minutes = m
	return nil
}

// ParseString parses a string
func (t *Time) ParseString(s string) error {
	// strings are in the form of HH:MM or
	// H:MM
	if len(s) < 4 || len(s) > 5 {
		return ErrInvalidString
	}
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return ErrInvalidString
	}
	h, err := strconv.ParseUint(parts[0], 10, strconv.IntSize)
	if err != nil {
		return ErrInvalidString
	}
	m, err := strconv.ParseUint(parts[1], 10, strconv.IntSize)
	if err != nil {
		return ErrInvalidString
	}

	return t.SetTime(uint(h), uint(m))
}

// Hours returns the hour part of the Time
func (t Time) Hours() uint {
	return t.hours
}

// SetHours sets the hour part of the Time
func (t *Time) SetHours(h uint) error {
	return t.SetTime(h, t.minutes)
}

// Minutes returns the minutes part of the Time
func (t Time) Minutes() uint {
	return t.minutes
}

// SetMinutes sets the minutes part of the Time
func (t *Time) SetMinutes(m uint) error {
	return t.SetTime(t.hours, m)
}

// String returns a string formatted as HH:MM
func (t Time) String() string {
	return fmt.Sprintf("%0d:%0d", t.hours, t.minutes)
}

// MarshalJSON marshals the Time value as JSON
func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

// UnmarshalJSON unmarshals JSON into a TIme value
func (t *Time) UnmarshalJSON(b []byte) error {
	var s string

	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	return t.ParseString(s)
}
