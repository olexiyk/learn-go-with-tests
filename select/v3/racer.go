package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(url1, url2 string) (string, error) {
	finishChan := make(chan string)
	go ping(url1, finishChan)
	go ping(url2, finishChan)
	return <-finishChan, nil
}

func ping(url string, c chan string) {
	http.Get(url)
	c <- url
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	finishChan := make(chan string)
	go ping(a, finishChan)
	go ping(b, finishChan)
	select {
	case x := <-finishChan:
		return x, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}
