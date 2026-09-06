package tokenbucket

import (
	"example.com/ratelimiter/internal/limiterclock"
)

type tokenbucket struct {
	BucketCapacity      int
	RefillRatePerSecond int

	Clock limiterclock.LimiterClock
}

type limiterOption func(*tokenbucket)

func WithClock(clock limiterclock.LimiterClock) limiterOption {
	return func(f *tokenbucket) {
		f.Clock = clock
	}
}

func NewFixedWindowCounter(BucketCapacity, RefillRatePerSecond int, opts ...limiterOption) *tokenbucket {
	f := &tokenbucket{
		BucketCapacity:      BucketCapacity,
		RefillRatePerSecond: RefillRatePerSecond,
		Clock:               &limiterclock.SystemClock{},
	}

	for _, option := range opts {
		option(f)
	}

	return f
}
