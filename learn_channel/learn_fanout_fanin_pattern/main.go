package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("In Fanout/FanIn Test code")
	input := gen()
	//fan-out : same input channel to multiple function
	c0 := factorial(input)
	c1 := factorial(input)
	c2 := factorial(input)
	c3 := factorial(input)
	c4 := factorial(input)
	c5 := factorial(input)
	c6 := factorial(input)
	c7 := factorial(input)
	c8 := factorial(input)
	c9 := factorial(input)
	//merge is fan-in. Multiple channels finally merged to 1
	for n := range merge(c0,c1,c2,c3,c4,c5,c6,c7,c8,c9){
		fmt.Println(n)
	}
}

func gen() chan int {
	c := make(chan int)
	go func() {
		//Generating 1000 factorial computation problem
		for i := 0; i < 100; i++ {
			for j := 3; j <= 13; j++ {
				c <- j
			}
		}
		close(c)
	}()
	return c
}

func factorial(c chan int) chan int {
	out := make(chan int)
	go func() {
		for cc := range c{
			out <- fact(cc)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	j := 1
	for i:=1;i<=n;i++{
		j*=i
	}
	return j
}

func merge(ch ...chan int) chan int{
	out := make(chan int)
    var wg sync.WaitGroup

	output := func(c chan int) {
		for n := range c{
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(ch))
	for _,c := range ch{
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
