package racer

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimer = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimer)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
    select {
    case <-ping(a):
        return a, nil
    case <-ping(b):
        return b, nil
    case <-time.After(timeout):
        return "", fmt.Errorf("timeoud out waiting for %s and %s", a, b)
    }
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
