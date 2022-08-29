package checkpoint

import "time"

const (
	BlockingPolicy = "blocking"
	WaitingPolicy  = "waiting"
)

// RateLimiterConfig defines the different limits that can be used across the rate limiter.
type RateLimiterConfig struct {
	Limits []Limit
}

// Limit is the definition of one limit used by the rate limiter.
// multiple limits can be used in the same time and will all combinate.
// the associated for each policy defined the behaviour that will be applied if this limit is reached.
// If one of the limit is reached with the blocking policy, it will be prioritized compared to waiting ones.
type Limit struct {
	Window   time.Duration
	Quantity int
	Policy   string
}

type RateLimiter struct {
	config *RateLimiterConfig
}

// NewRateLimiter returns a RateLimiter pointer ater ensuring all the fields are correctly set and initialized.
func NewRateLimiter(cfg *RateLimiterConfig) (*RateLimiter, error) {
	return nil, nil
}
