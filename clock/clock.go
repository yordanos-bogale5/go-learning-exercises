// Package clock manages adding and subtracting times to a clock
package clock

import "fmt"

// BenchmarkAddMinutes-4           100000000               16.3 ns/op             0 B/op          0 allocs/op
// BenchmarkSubtractMinutes-4      100000000               16.2 ns/op             0 B/op          0 allocs/op
// BenchmarkCreateClocks-4         200000000                7.52 ns/op            0 B/op          0 allocs/op

// Clock stores the hours and minutes of a clock with minutes and hours rolled over
// I prefer hour and minutes even though the logic is lengthier.
// Run time is the same as using a Clock wiht only minutes
type Clock struct {
	h int
	m int
}

// New returns a clock of type Clock which represents hours and minutes.
// It accounts for overflow in hours and minutes along with negative hours
// and minutes
func New(h int, m int) Clock {
	extraHours, minutes := m/60, m%60

	// account for negative number for minutes
	if minutes < 0 {
		extraHours--
		minutes += 60
	}

	hours := (h + extraHours) % 24

	// account for negative hours
	if hours < 0 {
		hours += 24
	}
	return Clock{h: hours, m: minutes}
}

// Add adds a minutes to a clock c of type Clock and returns a new Clock
func (c Clock) Add(a int) Clock {
	return New(c.h, c.m+a)
}

// Subtract subtracts a minutes to a clock c of type Clock and returns a new Clock
func (c Clock) Subtract(a int) Clock {
	return New(c.h, c.m-a)
}

// String converts the clock to a string of format hh:mm
func (c *Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
