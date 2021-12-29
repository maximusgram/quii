package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
    aDuration := measureResponseTime(a)
    bDuration := measureResponseTime(b)

    if aDuration < bDuration {
        return a
    }

    return b
}

func measureResponseTime(a string) time.Duration {
    startA := time.Now()
    http.Get(a)
    return time.Since(startA)
}