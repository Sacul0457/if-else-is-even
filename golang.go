package main

import (
	"fmt"
	"sync"
	"time"
)

type IsEvenOperator struct {
	checked  uint
	checkers uint
	mu       sync.Mutex
	timezone time.Location
}

func New(timezone time.Location) IsEvenOperator {
	return IsEvenOperator{timezone: timezone}
}

func (c *IsEvenOperator) GetCheckers() uint {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.checkers
}

func (c *IsEvenOperator) GetChecked() uint {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.checked
}

func (c *IsEvenOperator) updateCheckers() {
	c.mu.Lock()
	c.checkers += 1
	c.mu.Unlock()
}

func (c *IsEvenOperator) updateChecked() {
	c.mu.Lock()
	c.checked += 1
	c.mu.Unlock()
}

func (c *IsEvenOperator) GetTime() time.Time {
	return time.Now().In(&c.timezone)
}

func (c *IsEvenOperator) NewChecker() func(a uint, r *bool) {
	c.updateCheckers()
	f := func(a uint, r *bool) {
		go c.processCheck(a, r)
	}
	return f
}

func (c *IsEvenOperator) CheckingSocket(in <-chan uint) <-chan bool {
	out := make(chan bool)
	go func() {
		defer close(out)
		for m := range in {
			var result bool
			c.processCheck(m, &result)
			out <- result
		}
	}()
	return out
}

func (c *IsEvenOperator) processCheck(a uint, r *bool) {
	c.updateChecked()
	*r = (a%2 == 0)
}

func main() {
	// Instructions:

	// Step №1: Create Operator
	loc, _ := time.LoadLocation("Europe/Moscow")
	operator := New(*loc)

	// Step №2: Create Checker
	checker := operator.NewChecker()

	// Step №3: Check.
	var a uint = 5
	var result bool

	checker(a, &result)

	// Boom!
	// You have just mastered is-even checking skill.

	// Step №4 (Very Advanced): Checking Socket
	// You will now learn production-like is-even checking

	// Step №4.1: Create Channel
	ch := make(chan uint)

	// Step №4.2: Create Socket
	sock := operator.CheckingSocket(ch)

	// Step №4.4(Optional): Listen to checks
	go func() {
		for b := range sock {
			fmt.Print("Is the value even? Answer: ", b, "\n")
		}
	}()

	// Step №4.4: Check!
	for i := range 10 {
		ch <- uint(i)
	}

	// Step №5: Get some stats
	fmt.Print("Checkers amount: ", operator.GetCheckers(), "\n")
	fmt.Print("Checks made: ", operator.GetChecked(), "\n")
	fmt.Print("Current time: ", operator.GetTime().String(), "\n")

	close(ch)
}
