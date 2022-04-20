package main

import "fmt"

/*
	A protection proxy is the kind of proxy which performs access control.
	It basically tries to check whether or not the object that is trying to proxy is actually allowed to be
	accessed.

	Let's suppose that we are simulating the process of cars and other vehicles being driven
*/

type Driven interface {
	Drive()
}

type Car struct{}

func (c Car) Drive() {
	fmt.Println("Car is being driven")
}

/*
	Imagine that you want the car to only be driven if you have a driver,
	and that driver is actually old enough.

	So what we could do then is to build a protection proxy on top of the car.
	So, once again, you would be reusing the car somehow, but would also be specifying the driver.
*/

type Driver struct {
	Age int
}

type CarProxy struct {
	car    Car
	driver *Driver
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 16 {
		c.car.Drive()
	} else {
		fmt.Println("Driver too young!")
	}
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{Car{}, driver}
}

func main() {
	car := NewCarProxy(&Driver{12})
	car.Drive()
}

/*
	This example shows a very common kind of pattern where, first of all, we are starting out with
	an object or some sort of struct which can be used as it is without any verification.
	But subsequently, we want to have additional verifications being made
	whenever someone uses this struct.
*/
