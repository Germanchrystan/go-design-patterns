package main

import "fmt"

/*
	We can take a look into an academic example that we would likely find
	in many books and online examples, the idea of states being represented
	by structures and having some sort of replacement of one structure with another.

	This demo may be rather confusing. We will see the classic implementation
	of a simple state machine with just two states.

	Let's imagine a light, that has two state: on or off.
*/

type Switch struct {
	State State
}

// We need an interface for the state
type State interface {
	On(sw *Switch)
	Off(sw *Switch)
}

/*
	This may seem like heavy overengineering.
	Why can't we just keep everything inside us?
	Why should we have the state?

	The idea is that the switch stays the same, but we use implementators
	of this interface and we switch from one implementator to another.
	When we switch the light on or off, we replace the value or the property 'State' within switch.
*/
type BaseState struct {
}

func (b *BaseState) On(sw *Switch) {
	fmt.Println("Light is already on")
}

func (b *BaseState) Off(sw *Switch) {
	fmt.Println("Light is already off")
}

/*
	This looks very confusing, why are we defining the base state, which also makes an assumption that
	we haven't really switched the state from one to another.
	The idea is that we really define on and off state subsequently as separate structs.
*/
//========================================================//
type OnState struct {
	BaseState
}

func NewOnState() *OnState {
	fmt.Println("Light turned on")
	return &OnState{BaseState{}}
}

func (o *OnState) Off(sw *Switch) {
	fmt.Println("Turning the light off")
	sw.State = NewOffState()
}

//========================================================//
type OffState struct {
	BaseState
}

func NewOffState() *OffState {
	fmt.Println("Light turned off")
	return &OffState{BaseState{}}
}

func (o *OffState) On(sw *Switch) {
	fmt.Println("Turning light on")
	sw.State = NewOnState()
}

//========================================================//
func (s *Switch) On() {
	s.State.On(s)
}

func (s *Switch) Off() {
	s.State.Off(s)
}

/*
	What is happening here is effectively double dispatches, the kind of that
	appear on the visitor design pattern, except that unlike the visitor design pattern, what happens
	here is comp´letely unnecessary, meaning that this entire model can be simplified
	to be much more effective.
*/
//========================================================//
func NewSwitch() *Switch {
	return &Switch{NewOffState()}
}

//========================================================//
func main() {
	sw := NewSwitch()
	sw.On()
	sw.Off()
	sw.Off() /*
		The state is already an off state.
		If you look at the off state, it does not have an off method.
		It has only an on method.
		So, by invoking the off method a second time, we are calling the off method of the base state.
		The base state method prints a string saying that the light is already off.

		This is a purely academic example. This is not the kind of implementation we are
		likely going to be building in the real world. But is is a good illustration.
	*/
}
