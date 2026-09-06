package slidingwindowlog

import (
	"sync"
	"time"

	"example.com/ratelimiter/internal/limiterclock"
)

type slidingwindowlog struct {
	WindowDuration time.Duration
	MaxRequests    int
	userCache      map[string][]time.Time
	mu             sync.Mutex

	Clock limiterclock.LimiterClock
}

type limiterOption func(*slidingwindowlog)

func WithClock(clock limiterclock.LimiterClock) limiterOption {
	return func(s *slidingwindowlog) {
		s.Clock = clock
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

// ConcurrentSafe
func (s *slidingwindowlog) AllowRequest(userId string) bool {

	s.mu.Lock()
	defer s.mu.Unlock()

	//start from the back of the slice of request times and count the number in the time window
	requests, requestcount, now := s.userCache[userId], 0, s.Clock.Now()
	for i := len(requests) - 1; i >= 0; i-- {
		if now.Sub(requests[i]) > s.WindowDuration {
			//prune requests - don't store old timestamps
			requests = requests[i+1:]
			break
		}
		requestcount++
	}

	if requestcount < s.MaxRequests {
		requests = append(requests, now)
		s.userCache[userId] = requests
		return true
	}

	return false
}

// ConcurrentSafe
// func (f *slidingwindowlog) AllowNRequests(userId string, requests int) bool {

// 	return false
// }
