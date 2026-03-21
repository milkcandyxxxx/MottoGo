package middleware

import (
	"golang.org/x/time/rate"
	"sync"
)

// Ratelimit 限流器
type Ratelimit struct {
	// key 是 IP，value 是 *rate.Limiter (官方限流器对象)
	limiters sync.Map
}

// Limit 限制每秒最多访问10次
func (r *Ratelimit) Limit(ip string) bool {

	l, _ := r.limiters.LoadOrStore(ip, rate.NewLimiter(rate.Limit(10), 10))

	limiter := l.(*rate.Limiter)
	return limiter.Allow()
}
