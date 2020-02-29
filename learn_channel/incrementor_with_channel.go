package main

import "fmt"

//run command : go run main.go incrementor_with_channel.go

func testIncrementorWithChannel() {

	c := inc(2)

	count := 0
	for x := range c {
		count++
		fmt.Println(x)
	}

	fmt.Println("Total count ", count)
}

func inc(n int) chan string {
	out := make(chan string)
	done := make(chan bool)
	for i:=0 ;i<n ;i++{
		go func() {
			for j:=0;j<20;j++{
				out <- "I am inc"
			}
			done <- true
		}()
	}

	go func() {
		for i:=0 ;i<n ;i++{
			<-done
		}
		close(out)
	}()
	return out
}
