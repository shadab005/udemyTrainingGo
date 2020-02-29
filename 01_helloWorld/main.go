package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	Value int
}

func printCounter(c Counter) {
	fmt.Printf("%d ", c.Value)
}

func Add10Alfa(c *Counter) {
	c.Value += 10
}

func Add10Bravo(c Counter) {
	c.Value += 10
}

func main() {
	var c Counter

	Add10Alfa(&c)
	printCounter(c)

	Add10Bravo(c)
	printCounter(c)
}

func count(str string)  {

	for i:=0;i< 10; i++ {
		fmt.Println("from ", str, " i = ", i)
		time.Sleep(time.Microsecond*500)
	}
}

func prіntItems() {
	var o sync.WaitGroup
	itemѕ := []string{"one", "two", "three"}
	for _, item := range itemѕ {
		o.Add(1)
		go func() {
			defer o.Done()
			fmt.Println(item)
		}()
	}
	o.Wait()
}

func function() bool {
	var c map[string]string
	if c["dog"] == "" {
		return true
	}
	c["dog"] = "Buster"
	return false
}