package system

import (
	"time"
)

// Sleep Sleeps the current process
func (plugin *OSPlugin) Sleep(number int, unit string) {
	var timeUnit time.Duration
	switch unit {
	case "MS":
		timeUnit = time.Millisecond
	case "S":
		timeUnit = time.Second
	default:
		timeUnit = time.Second
	}

	time.Sleep(time.Duration(number) * timeUnit)
}
