package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(even, odd, quit)

	receive(even, odd, quit)

	fmt.Println("Exiting Main")
}

func receive(even, odd, quit <-chan int) {
	for {
		select {
		case v := <-quit:
			fmt.Println("From quit channel", v)
			return
		case v := <-even:
			fmt.Println("From even channel", v)
		case v := <-odd:
			fmt.Println("From odd channel", v)
		}
	}

}

func send(even, odd, quit chan<- int) {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

	close(even)
	close(odd)
	quit <- 1
}
