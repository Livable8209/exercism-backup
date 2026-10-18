package clock

import "fmt"

const minutesPerDay int = 1440

// Define the Clock type here.
type Clock struct {
	// timestamp thats clamped to 0 - 1440, uses double modulo to do wrap around
	timestamp int
}

func New(h, m int) Clock {
	newTimestamp := normalizeTime(h*60 + m)
	return Clock{
		timestamp: newTimestamp,
	}
}

func (c Clock) Add(m int) Clock {
	addedTimestamp := normalizeTime(c.timestamp + m)
	return Clock{
		timestamp: addedTimestamp,
	}
}

func (c Clock) Subtract(m int) Clock {
	subtractedTimestamp := normalizeTime(c.timestamp - m)
	return Clock{
		timestamp: subtractedTimestamp,
	}
}

func (c Clock) String() string {
	hours := c.timestamp / 60
	minutes := c.timestamp % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

func normalizeTime(minutes int) int {
	return ((minutes % minutesPerDay) + minutesPerDay) % minutesPerDay
}
