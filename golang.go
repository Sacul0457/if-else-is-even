package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

type IsEvenOperator struct {
	checked  uint
	checkers uint
	mu       sync.Mutex
	timezone *time.Location
}

func New(timezone *time.Location) IsEvenOperator {
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
	return time.Now().In(c.timezone)
}

func (c *IsEvenOperator) NewChecker() func(a int, r *bool) {
	c.updateCheckers()
	return func(a int, r *bool) {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.processCheck(a, r)
		}()
		wg.Wait()
	}
}

func (c *IsEvenOperator) CheckingSocket(in <-chan int) <-chan bool {
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

func (c *IsEvenOperator) processCheck(a int, r *bool) {
	c.updateChecked()
	*r = (a%2 == 0)
}

//		 // Instructions:

//		 // Step №1: Create Operator
//		 loc, _ := time.LoadLocation("Europe/Moscow")
//		 operator := New(loc)

//		 // Step №2: Create Checker
//		 checker := operator.NewChecker()

//		 // Step №3: Check.
//		 var a uint = 5
//		 var result bool

//		 checker(a, &result)
//		 fmt.Print("Is ", a, " even? ", result, "\n")

//		 // Boom!
//		 // You have just mastered is-even checking skill.

//		 // Step №4 (Very Advanced): Checking Socket
//		 // You will now learn production-like is-even checking

//		 // Step №4.1: Create Channel
//		 ch := make(chan uint)

//		 // Step №4.2: Create Socket
//		 sock := operator.CheckingSocket(ch)

//		 // Step №4.3: Listen to checks
//		 var wg sync.WaitGroup
//		 wg.Add(1)
//		 go func() {
//		 	defer wg.Done()
//		 	for b := range sock {
//		 		fmt.Print("Is the value even? Answer: ", b, "\n")
//		 	}
//		 }()

//		 // Step №4.4: Check!
//		 for i := range 10 {
//		 	ch <- uint(i)
//		 }
//		 close(ch)

//		 // Step №5: Get some stats
//		 fmt.Print("Checkers amount: ", operator.GetCheckers(), "\n")
//		 fmt.Print("Checks made: ", operator.GetChecked(), "\n")
//		 fmt.Print("Current time: ", operator.GetTime().String(), "\n")

func main() {
	loc, _ := time.LoadLocation("Europe/Moscow")
	op := New(loc)
	checker := op.NewChecker()

	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Only one argument must be provided, received: %s", args)
		return
	}
	str := args[1]
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Invalid integer provided: %s", str)
		return
	}

	var result bool
	checker(num, &result)

	resultStr := "even"
	if !result {
		resultStr = "odd"
	}

	fmt.Printf("The number %d is %s", num, resultStr)
	return
}
