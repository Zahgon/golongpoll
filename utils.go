package golongpoll

import (
	"time"
)

// millisecondStringToTime converts a string representation of milliseconds
// into the corresponding Time value in UTC.  This is used to convert
// timestamps sent by browsers that use javascript's Date.getTime() and then
// pass that number as a url param (hence why it winds up a string).
// IF an invalid string input is used, then time will default to Time{} and the
// error return value will be non-nil.
func millisecondStringToTime(ms string) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// timeToEpochMilliseconds converts a Time type to the corresponding
// number of milliseconds since epoch (Jan 1 1970) in UTC.  Note if the Time
// input is before 1970, the corresponding milliseconds value is negative.
func timeToEpochMilliseconds(t time.Time) int64 { _ = "STUB: not implemented"; return 0 }
