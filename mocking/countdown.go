package mocking

import (
	"fmt"
	"io"
)

type Sleeper interface {
	Sleep()
}

const (
	countdownStart = 3
	finalWord      = "Go!"
)

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprint(
		w,
		finalWord,
	)
}
