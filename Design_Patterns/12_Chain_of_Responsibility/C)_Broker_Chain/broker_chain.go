package main

import (
	"fmt"
	"sync"
)

/*
	We are now going to take a look at a much more advanced implementation of the chain of responsibility
	design pattern. This design pattern is going to combine multiple approaches.

	We are going to have the mediator design pattern, a centralized component
	everyone talks to.

	We are going to have the observer design pattern, because this mediator is going to be
	observable.

	There is going to be lots of other enterprise approaches, including the idea of command query separation.
*/
//=================================================//
const (
	Attack Argument = iota
	Defense
)

type Argument int

type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}

//=================================================//
//Observer Design Pattern
type Observer interface {
	Handle(q *Query)
}

type Observable interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	Fire(q *Query)
}

//=================================================//
type Game struct {
	observers sync.Map /* This is going to allow us to keep a map
	of every single subscriber,
	to iterate this map to go through every single subscriber,
	and notify on that subscriber.
	*/
}

func (g Game) Subscribe(o Observer) {
	// Adding observer to list
	g.observers.Store(o, struct{}{})
}

func (g Game) Unsubscribe(o Observer) {
	g.observers.Delete(o)
}

func (g Game) Fire(q *Query) {
	g.observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).Handle(q)
		return true
	})
}

//=================================================//

type Creature struct {
	Name            string
	game            *Game
	attack, defense int
	/*
		In the previous example, we only applied the modifiers explicitly by calling
		the handle method.
		What we want to be able to do now is to apply these modifiers automatically.
		As soon as we make a modifier, passing it to a creature, it automatically gets applied.
		The attack and defense values here only store the initial values, not the calculated ones.
	*/
}

func NewCreature(game *Game, name string, attack int, defense int) *Creature {
	return &Creature{game: game, Name: name, attack: attack, defense: defense}
}

func (c *Creature) Attack() int {
	q := Query{c.Name, Attack, c.attack}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature) Defense() int {
	q := Query{c.Name, Defense, c.defense}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s(%d/%d)", c.Name, c.Attack(), c.Defense())
}

//=================================================//
type CreatureModifier struct {
	game     *Game
	creature *Creature
}

func (c *CreatureModifier) Handle(q *Query) {
	// empty
}

//-------------------------------------//
type DoubleAttackModifier struct {
	CreatureModifier
}

func NewDoubleAttackModifier(g *Game, c *Creature) *DoubleAttackModifier {
	d := &DoubleAttackModifier{CreatureModifier{g, c}}
	g.Subscribe(d)

	return d
}

func (d *DoubleAttackModifier) Handle(q *Query) {
	if q.CreatureName == d.creature.Name &&
		q.WhatToQuery == Attack {
		q.Value *= 2
	}
}

func (d *DoubleAttackModifier) Close() error {
	d.game.Unsubscribe(d)
	return nil
}

//=================================================//
func main() {
	game := &Game{sync.Map{}}
	goblin := NewCreature(game, "Strong Goblin", 2, 2)
	fmt.Println(goblin.String())
	// Here we can apply the double attack modifier temporarily
	{
		m := NewDoubleAttackModifier(game, goblin)
		fmt.Println(goblin.String())
		m.Close()
	}
	fmt.Println(goblin.String())

}
