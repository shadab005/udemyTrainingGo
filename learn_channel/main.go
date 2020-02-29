package main

import (
	"fmt"
	"time"
)

func main() {
	testIncrementorWithChannel()
}


//N producer 1 consumer
func test_N_to_1_SemaphoreParadigm() {

	c := make(chan int)
	done := make(chan bool)

	n := 10

	//Running n go routines concurrently
	for i := 0; i < n; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- j
			}
			done <- true
		}()
	}

	//Go routine to wait for above go routines to complete
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(c)
	}()

	for data := range c {
		fmt.Println(data)
	}

}

//1 producer N consumer
func test_1_to_N_SemaphoreParadigm() {
	c := make(chan int)
	done := make(chan bool)
	n := 10

	//One producer
	go func() {
		for j := 0; j < 1000; j++ {
			c <- j
		}
		close(c)
	}()

	//n consumer
	for i := 0; i < n; i++ {
		time.Sleep(time.Millisecond)
		go func() {
			for x := range c {
				fmt.Println("Consumer ", i, " Value ", x)
			}
			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}

	fmt.Println("Done with 1 to N")

}

func testChannelAsArgumentAndReturn() {

	c := incrementor()
	p := puller(c)

	for x := range p {
		fmt.Println("Sum = ", x)
	}

	fmt.Println("Done with testChannelAsArgumentAndReturn")

}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		sum := 0
		for n := range c {
			fmt.Println(n)
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func testCreateDeadlock() {
	c1 := make(chan int)
	c1 <- 1 //put this in a go routine. This will prevent deadlock then
	fmt.Println(<-c1)
}

// NOTE: There is also a concept of direction channel

func testPipeLinePattern() {
	/*c := generate(1,2,3,4)
	out := square(c)
	for n := range out {
		fmt.Println(n)
	}*/

	fmt.Println("Square")
	for n := range square(generate(1, 2, 3)) {
		fmt.Println(n)
	}

	fmt.Println("Square of Square")
	for n := range square(square(generate(1, 2, 3))) {
		fmt.Println(n)
	}

}

func generate(nums ...int) chan int {
	c := make(chan int)
	go func() {
		for _, x := range nums {
			c <- x
		}
		close(c)
	}()
	return c
}

func square(c chan int) chan int {
	out := make(chan int)
	go func() {
		for num := range c {
			out <- num * num
		}
		close(out)
	}()
	return out
}
