package fixedwindowcounter

import (
	"sync"
	"time"

	"example.com/ratelimiter/internal/limiterclock"
)

type fixedwindowcounter struct {
	WindowDuration time.Duration
	MaxRequests    int
	userCache      map[string]int
	mu             sync.Mutex

	Clock limiterclock.LimiterClock
}

type limiterOption func(*fixedwindowcounter)

func WithClock(clock limiterclock.LimiterClock) limiterOption {
	return func(f *fixedwindowcounter) {
		f.Clock = clock
	}
}

func NewFixedWindowCounter(WindowDuration time.Duration, MaxRequests int, opts ...limiterOption) *fixedwindowcounter {
	f := &fixedwindowcounter{
		WindowDuration: WindowDuration,
		MaxRequests:    MaxRequests,
		userCache:      make(map[string]int),
		Clock:          &limiterclock.SystemClock{},
	}

	for _, option := range opts {
		option(f)
	}

	return f
}

// ConcurrentSafe
func (f *fixedwindowcounter) AllowRequest(userId string) bool {

	return false
}

// ConcurrentSafe
func (f *fixedwindowcounter) AllowNRequests(userId string, requests int) bool {

	return false
}
