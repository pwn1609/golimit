# rate-limiter-library

A small Go library implementing a few common rate-limiting algorithms.

## Why this exists

This is a weekend project, written entirely by hand, to exercise my manual coding muscles. Don't get too excited though, this README was obviously AI generated.

## Layout

- `limiter.go` — the shared `Limter` interface (`AllowRequest`, `AllowNRequests`)
- `internal/limterclock/` — `LimiterClock` time abstraction with `SystemClock`
  for real use and `FakeClock` for deterministic tests
- `fixedwindowcounter/` — fixed-window counter limiter and its
  functional-options constructor (`NewFixedWindowCounter`, `WithClock`)

## Design notes

- Every algorithm implements the common `Limter` interface so they're
  interchangeable.
- Time is injected via `LimiterClock` rather than calling `time.Now()`
  directly, so tests can drive the clock with `FakeClock`.
- Limiters are constructed with the functional options pattern, e.g. passing
  `WithClock` to override the default `SystemClock`.

## Usage

```go
import (
	"time"

	"example.com/ratelimiter/fixedwindowcounter"
)

// Allow at most 100 requests per minute.
limiter := fixedwindowcounter.NewFixedWindowCounter(time.Minute, 100)
```

The API is unstable; expect signatures to move around as more algorithms land.
