package carrot

import (
	"time"
)

type Base struct {
	URL, Proto string
	Count      int
	Delay      int
	TickDelay  int
	Path       string
}

type Routine struct {
	SendTime    time.Time
	ReceiveTime time.Time
	Diff        time.Duration // milliseconds
	ReceivedMsg string
}
