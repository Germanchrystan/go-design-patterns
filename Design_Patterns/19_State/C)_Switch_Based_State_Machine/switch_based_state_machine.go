package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	There is yet another type of state machine to show,
	and this state machine is very special, because instead
	of having a map of the different transitions, what happens is that
	the information is encoded somewhere else.
	In this case, it is encoded inside a switch statement.

	First of all, let's set up an scenario.
	We are going to try to model a combination lock.
	It consists on basically four digits for the lock, and somebody makes up
	the combination to unlock it. If we enter the wrong combination the lock remains locked.
*/

type State int

const (
	Locked State = iota
	Failed
	Unlocked
)

func main() {
	code := "1234"
	state := Locked
	entry := strings.Builder{}
	for {
		switch state {
		case Locked:
			r, _, _ := bufio.NewReader(os.Stdin).ReadRune()
			entry.WriteRune(r)

			if entry.String() == code {
				state = Unlocked
				break
			}

			// As soon as we enter an incorrect digit, the lock can detect it.
			if strings.Index(code, entry.String()) != 0 {
				state = Failed
			}
		case Failed:
			fmt.Println("FAILED")
			entry.Reset()
			state = Locked
		case Unlocked:
			fmt.Println("UNLOCKED")
			return
		}
	}
}
