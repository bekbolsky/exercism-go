package clock

import (
	"fmt"
)

// Clock struct represents a clock with hour and minute
type Clock struct {
	hour, minute int
}

// New creates a new clock with hour and minute
func New(h, m int) Clock {
	h = (h + m/60) % 24
	m = m % 60
	if m < 0 {
		m += 60
		h--
	}
	if h < 0 {
		h += 24
	}
	if h > 23 {
		h -= 24
	}
	if m > 59 {
		m -= 60
		h++
	}
	return Clock{h, m}
}

// Add increases the clock by minutes
func (c Clock) Add(m int) Clock {
	return New(c.hour, c.minute+m)
}

// Subtract decreases the clock by minutes
func (c Clock) Subtract(m int) Clock {
	return New(c.hour, c.minute-m)
}

// String returns the string representation of the clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
