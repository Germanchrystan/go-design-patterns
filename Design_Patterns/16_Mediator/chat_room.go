package main

import (
	"fmt"
)

// Let's simulate a chat room application

type Person struct {
	Name string
	// Pointer to the mediator
	Room *ChatRoom
	// Log of all messages sent or received
	chatLog []string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (p *Person) Receive(sender, message string) {
	s := fmt.Sprintf("%s:%s\n", sender, message)
	fmt.Printf("[%s chat session]: %s", p.Name, s)
	p.chatLog = append(p.chatLog, s)
}

// Sending something to the general room
func (p *Person) Say(message string) {
	p.Room.Broadcast(p.Name, message)
}

func (p *Person) PrivateMessage(who, message string) {
	p.Room.Message(p.Name, who, message)
}

// Let's define the actual room
// The room is just a collection of people
// This is the mediator. People in the chat don't have a pointer to one another
type ChatRoom struct {
	people []*Person
}

func (c *ChatRoom) Broadcast(source, message string) {
	for _, p := range c.people {
		if p.Name != source {
			p.Receive(source, message)
		}
	}
}

func (c *ChatRoom) Message(src, dst, msg string) {
	for _, p := range c.people {
		// If the person is found, command them to receive the message
		if p.Name == dst {
			p.Receive(src, msg)
		}
	}
}

func (c *ChatRoom) Join(p *Person) {
	joinMsg := p.Name + " joins the chat"
	c.Broadcast("Room", joinMsg)
	p.Room = c
	c.people = append(c.people, p)
}

func main() {
	room := ChatRoom{}
	john := NewPerson("John")
	jane := NewPerson("Jane")

	room.Join(john)
	room.Join(jane)

	john.Say("hi room")
	jane.Say("Hi, John")

	simon := NewPerson("Simon")
	room.Join(simon)
	simon.Say("hi everyone!")

	jane.PrivateMessage("Simon", "Glad you could join us!")
}
