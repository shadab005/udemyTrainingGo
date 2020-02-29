package main

import (
	"fmt"
	"time"
)

func main() {
	testChannel()
}

func testChannel() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Sending data to channel ", i)
			c <- i
			fmt.Println("Channel is empty now")
		}
	}()

	go func() {
		for {
			fmt.Println("Receiving data from channel")
			fmt.Println("Data received from channel : ", <-c)
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}
