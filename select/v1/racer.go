package racer

import (
	"net/http"
	"time"
)

func Racer(url1, url2 string) string {
	t1 := measureResponseTime(url1)
	t2 := measureResponseTime(url2)
	if t1 > t2 {
		return url2
	}
	return url1
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
