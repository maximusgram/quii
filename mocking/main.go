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

type DefaultSleeper struct {}

func (d *DefaultSleeper) Sleep() {
    time.Sleep(1 * time.Second)
}

const (
    finalWord = "Go!"
    countdownStart = 3
)

func Countdown(out io.Writer, sleep Sleeper) {
    for i := countdownStart; i >= 1; i-- {
        sleep.Sleep()
        fmt.Fprintln(out, i)
    }
    fmt.Fprint(out, "Go!")
}
func main() {
    Countdown(os.Stdout, &DefaultSleeper{})
}