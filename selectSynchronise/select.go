package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	// `select` lets you wait on multiple channels. The first one to write to its channel "wins" and its code gets executed first.
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout): // time.After will return a channel after the specified duration has elapsed
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// func measureResponseTime(url string) time.Duration {
// 	start := time.Now()
// 	http.Get(url)
// 	return time.Since(start)
// }

// this will create a new goroutine for each ping. It will write to its channel once http.Get resolves.
func ping(url string) chan struct{} {
	ch := make(chan struct{}) // need to `make` channels because initializing via `var` will initialize to zero value of nil
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
