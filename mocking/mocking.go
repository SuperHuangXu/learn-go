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

type ConfigurableSleeper struct {
  duration time.Duration
}

func (c ConfigurableSleeper) Sleep() {
  time.Sleep(c.duration)
}

const finalWord = "Go!"

const countdownStart = 3

func Countdown(out io.Writer, sleeper Sleeper) {
  for i := countdownStart; i > 0; i-- {
    sleeper.Sleep()
    fmt.Fprintln(out, i)
  }

  sleeper.Sleep()
  fmt.Fprint(out, finalWord)
}

func main() {
  sleeper := ConfigurableSleeper{1 * time.Second}
  Countdown(os.Stdout, sleeper)
}
