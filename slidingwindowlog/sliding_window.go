package slidingwindowlog

import (
	"time"

	"example.com/ratelimiter/internal/limiterclock"
)

type slidingwindowlog struct {
	WindowDuration time.Duration
	MaxRequests    int

	Clock limiterclock.LimiterClock
}

type limiterOption func(*slidingwindowlog)

func WithClock(clock limiterclock.LimiterClock) limiterOption {
	return func(f *slidingwindowlog) {
		f.Clock = clock
	}
}

func NewSlidingWindowLog(WindowDuration time.Duration, MaxRequests int, opts ...limiterOption) *slidingwindowlog {
	f := &slidingwindowlog{
		WindowDuration: WindowDuration,
		MaxRequests:    MaxRequests,
		Clock:          &limiterclock.SystemClock{},
	}

	for _, option := range opts {
		option(f)
	}

	return f
}
