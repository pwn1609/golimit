package fixedwindowcounter

import (
	"sync"
	"time"

	"example.com/ratelimiter/internal/limiterclock"
)

type fixedwindowcounter struct {
	WindowDuration     time.Duration
	MaxRequests        int
	userCache          map[string]int
	mu                 sync.Mutex
	currentWindowStart time.Time

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

	f.currentWindowStart = f.Clock.Now()

	for _, option := range opts {
		option(f)
	}

	return f
}

// ConcurrentSafe
func (f *fixedwindowcounter) AllowRequest(userId string) bool {

	//check if the current window time has elapsed.
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.Clock.Now().Sub(f.currentWindowStart) > f.WindowDuration {
		//Yes - set the currentWindowStart
		f.currentWindowStart = f.Clock.Now()
		f.userCache = make(map[string]int)
	}
	//lookup UserID in map for number of requests in this window.
	if f.userCache[userId] < f.MaxRequests {
		f.userCache[userId]++
		return true
	}

	return false
}

// ConcurrentSafe
// func (f *fixedwindowcounter) AllowNRequests(userId string, requests int) bool {

// 	return false
// }
