package limiterclock

import "time"

type LimiterClock interface {
	Now() time.Time
}

type SystemClock struct{}

func (s *SystemClock) Now() time.Time {
	return time.Now()
}

type FakeClock struct {
	time time.Time
}

func (f *FakeClock) Now() time.Time {
	return f.time
}

func (f *FakeClock) SetTime(time time.Time) {
	f.time = time
}
