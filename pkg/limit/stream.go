package limit

var UserLimiter = make(map[string]*RateLimiter)

func CheckFlowControl(uid string) bool {
	limiter := UserLimiter[uid]

	if limiter == nil {
		return false
	}
	return limiter.Limit()
}
