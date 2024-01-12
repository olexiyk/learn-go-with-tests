package racer

import (
	"net/http"
)

func Racer(url1, url2 string) string {
	select {
	case <-ping(url1):
		return url1
	case <-ping(url2):
		return url2
	}
}

func ping(url string) chan struct{} {
	pingChannel := make(chan struct{})
	go func(u string) {
		http.Get(u)
		close(pingChannel)
	}(url)

	return pingChannel
}
