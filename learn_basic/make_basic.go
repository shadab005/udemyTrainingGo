package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
			time.Sleep(5000 * time.Millisecond)
		}
	}()
	i := 0
	for {
		if i > 100 {
			fmt.Println("Breaking")
			break
		}
		select {
		case x := <-ch:
			fmt.Println("Here", x)
		case <-time.After(1000 * time.Millisecond):
			fmt.Println("Here in time After")

		}
		fmt.Println("Incrementing")
		i++
	}

}
