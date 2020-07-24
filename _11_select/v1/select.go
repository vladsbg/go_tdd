package server_racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	durationA := measureResponseTime(a)
	durationB := measureResponseTime(b)

	if durationA > durationB {
		return b
	}
	return a
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
