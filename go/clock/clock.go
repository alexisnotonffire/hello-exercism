// Package clock provides a clock
package clock

import "strconv"

// Clock is 24hr and has hours and minutes.
type Clock struct {
	hour   int
	minute int
}

// New take hours and inutes and turns them into a clock.
// Converting hours to minutes, totalling the minutes and using
// modulo division to provide a 24hr clock?
func New(h, m int) Clock {
	totalMinutes := h*60 + m
	m = totalMinutes % 60
	h = (totalMinutes - m) / 60 % 24

	if m < 0 {
		m += 60
		h--
	}
	if h < 0 {
		h += 24
	}

	return Clock{hour: h, minute: m}
}

// Add increments the provided minutes to a clock
func (c Clock) Add(m int) Clock {
	totalMinutes := c.hour*60 + c.minute + m
	m = totalMinutes % 60
	h := (totalMinutes - m) / 60 % 24

	if m < 0 {
		m += 60
		h--
	}
	if h < 0 {
		h += 24
	}

	return Clock{hour: h, minute: m}
}

// String displays a readable clockface
func (c Clock) String() string {
	var hours, minutes string
	if c.hour < 10 {
		hours = "0" + strconv.Itoa(c.hour)
	} else {
		hours = strconv.Itoa(c.hour)
	}
	if c.minute < 10 {
		minutes = "0" + strconv.Itoa(c.minute)
	} else {
		minutes = strconv.Itoa(c.minute)
	}

	return hours + ":" + minutes
}

// CompareClocks checks the clocks are set to the same time
func CompareClocks(a, b Clock) bool {
	return (a.hour == b.hour) && (a.minute == b.minute)
}
