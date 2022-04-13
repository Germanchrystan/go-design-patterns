package main

import "fmt"

/*
	One issue that developers face in different programming languages is how to combine the functionality of
	several structures inside a single structure.
	Unfortunately, with languages that do not support multiple inheritance, that becomes difficult.

	With Go there is no inheritance, there is only aggregation, and there are situations where aggregation
	doesnt't give us the desired result.

	Let's se where this can go wrong, as well as how this problem can be fixed.
*/
// Let's assume a Dragon is a mixture of a Bird and a Lizard
type Bird struct {
	Age int
}

func (b *Bird) Fly() {
	if b.Age >= 10 {
		fmt.Println("Flying!")
	}
}

type Lizard struct {
	Age int
}

func (l *Lizard) Crawl() {
	if l.Age < 10 {
		fmt.Println("Crawling!")
	}
}

type Dragon struct {
	Bird
	Lizard
} /*
	This aggregation can cause problems.
	In this example, both the bird and the lizard have an Age property
*/

func main() {
	d := Dragon{}
	//d.Age = 10  => This will cause an "ambiguous selector" error
	d.Bird.Age = 10
	d.Lizard.Age = 10
	d.Fly()
	d.Crawl()
	/*
		The problem with this scenario is that you can introduce inconsistency into the behavior
		of the dragon, by setting the different ages to different values.
		After all, you don't need to separate fields. It's a single age, we should keep it in a single field.
	*/
}

// So we could have getters and setters that fix this issue
func (d *Dragon) Age() int {
	return d.Bird.Age
}

func (d *Dragon) SetAge(age int) {
	d.Bird.Age = age
	d.Lizard.Age = age
}

func main2() {
	d := Dragon{}
	d.SetAge(5)
	/*
		Unfortunately, it is still very possible to make this object inconsistent by going into the
		dragon explicitly.
	*/
	d.Bird.Age = 6
	/*
		So there is no real solution to this, at least not in Go.
		There is no language feature that will allow you to regularize this whole situation.
		Whate we can try to do is to design the entire set of structs so that this is avoided and so that
		instead of simply aggregating, we will build a decorator around the bird and the lizard types.
	*/
}

// Let's start again
// First of all, we want an interface for objects with an age property
type Aged interface {
	Age() int
	SetAge(age int)
} /*
	Getters and Setters are not particularly idiomatic in Go.
	In most cases you want to avoid having getters and setters, but in this particular situation,
	there is no way around it.
*/

type BirdCorrected struct {
	age int // age is private here
}

func (b *BirdCorrected) Age() int       { return b.age }
func (b *BirdCorrected) SetAge(age int) { b.age = age }
func (b *BirdCorrected) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying!")
	}
}

type LizardCorrected struct {
	age int
}

func (l *LizardCorrected) Age() int       { return l.age }
func (l *LizardCorrected) SetAge(age int) { l.age = age }

func (l *LizardCorrected) Crawl() {
	if l.age <= 10 {
		fmt.Println("Crawling!")
	}
}

// Now we can construct the Dragon, but differently
type DragonCorrected struct {
	bird   BirdCorrected
	lizard LizardCorrected
}

func (d *DragonCorrected) SetAge(age int) {
	d.bird.SetAge(age)
	d.lizard.SetAge(age)
} /*
	Because the age property is private in BirdCorrected and LizardCorrected,
	values cannot be changed explicitly.
	Instead, we only operate on behaviours.
*/

func (d *DragonCorrected) Fly() {
	d.bird.Fly()
}

func (d *DragonCorrected) Crawl() {
	d.lizard.Crawl()
}

func NewDragon() *DragonCorrected {
	return &DragonCorrected{BirdCorrected{}, LizardCorrected{}}
}

func main3() {
	d := DragonCorrected{}
	d.SetAge(6)
	d.Fly()
	d.Crawl()
}

/*
	In the DragonCorrected class we have constructed a decorator,
	which extends the behaviors of the types Bird and Lizard.

	What it is doing really is providing better access to the underlying fields of both the bird
	and the lizard, because the age must be set consistently.
	In addition, DragonCorrected combines the behaviors of both the bird and the lizard by providing
	the interface members with the same names.

*/
