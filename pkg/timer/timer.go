package timer

import (
	"time"
)

var canceled = make(chan int)

func CancelTimer() {
	select {
	case canceled <- 0:
	default:
	}
}

func PlanAction(fn func(), seconds time.Duration) {
	// cancel all current timers
	select {
	case canceled <- 0:
	default:
	}

	go actionTimer(fn, seconds)
}

func actionTimer(fn func(), seconds time.Duration) {
	select {
	case <-time.After(seconds * time.Second):
		fn()
	case <-canceled:
	}
}
