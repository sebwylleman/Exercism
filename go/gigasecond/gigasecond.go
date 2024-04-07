// Package gigasecond uses seconds to express time intervals.A gigasecond is one thousand million seconds. That is a one with nine zeros after it.
package gigasecond

// import path for the time package from the standard library
import (
	"time"
)

// AddGigasecond returns a new time one gigasecond (10^9 seconds) after the provided timestamp 't'.
func AddGigasecond(t time.Time) time.Time {
	gs := time.Duration(1e9) * time.Second
	return t.Add(gs)
}
