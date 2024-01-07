package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

func Countdown(buffer io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(buffer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(buffer, "Go!")
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleepFn  func(time.Duration)
}

func (self ConfigurableSleeper) Sleep() {
	self.sleepFn(self.duration)
}

func main() {
	Countdown(os.Stdout, ConfigurableSleeper{time.Millisecond * 100, time.Sleep})
}
