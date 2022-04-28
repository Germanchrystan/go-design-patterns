package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	The states and the transitions should not be defined by some heavy weight constructs,
	like structures.
	We could define states by just defining constants.

	Let's simulate the states of a phone
*/

type State int

const (
	OffHook State = iota
	Connecting
	Connected
	OnHold
	OnHook
)

// There is one problem with Go, and that is the printing of this constants
// We have to use a string generator
func (s State) String() string {
	switch s {
	case OffHook:
		return "OffHook"
	case Connecting:
		return "Connecting"
	case Connected:
		return "Connected"
	case OnHold:
		return "OnHold"
	case OnHook:
		return "OnHook"
	}
	return "Unknown"
}

/*
	Now we must implement the triggers.
	The triggers are explicit definitions of what can cause us to go from
	one state to another.
*/

type Trigger int

const (
	CallDialed Trigger = iota
	HungUp
	CallConnected
	PlacedOnHold
	TakenOffHold
	LeftMessage
)

func (t Trigger) String() string {
	switch t {
	case CallDialed:
		return "CallDialed"
	case HungUp:
		return "HungUp"
	case CallConnected:
		return "CallConnected"
	case PlacedOnHold:
		return "PlacedOnHold"
	case TakenOffHold:
		return "TakenOffHold"
	case LeftMessage:
		return "LeftMessage"
	}
	return "Unknown"
}

/*
	We now need to define the rules which transition us from one state to another.
	For example, when we dial the call, we move from the OffHook state to the Connecting state.
	We should define our own map.

	First me define a TriggerResult struct.
	This is a combination of a trigger and the state you transition to when that trigger happens
*/

type TriggerResult struct {
	Trigger Trigger
	State   State
}

// From any given state, we might transition to more than one state, depending on the trigger
var rules = map[State][]TriggerResult{
	OffHook: {
		{CallDialed, Connecting},
	},
	Connecting: {
		{HungUp, OnHook},
		{CallConnected, Connected},
	},
	Connected: {
		{LeftMessage, OnHook},
		{HungUp, OnHook},
		{PlacedOnHold, OnHold},
	},
	OnHold: {
		{TakenOffHold, Connected},
		{HungUp, OnHook},
	},
}

func main() {
	state, exitState := OffHook, OnHook
	for ok := true; ok; ok = state != exitState {
		fmt.Println("The phone is currently", state) //We don't need to call the String method explicitly
		fmt.Println("Select a trigger:")

		for i := 0; i < len(rules[state]); i++ {
			tr := rules[state][i]
			fmt.Println(strconv.Itoa(i), ".", tr.Trigger)
		}

		input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		i, _ := strconv.Atoi(string(input))

		tr := rules[state][i]
		state = tr.State
	}
	fmt.Println("We are done using the phone")
}
