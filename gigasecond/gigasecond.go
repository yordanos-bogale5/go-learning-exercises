// Package gigasecond provides a function to add a gigasecond to a timestame
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// AddGigasecond cakes in the current time stamp and outputs the time
// 1 Gigasecond from that time
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1E9)
}
