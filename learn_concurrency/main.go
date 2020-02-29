package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	testWithAtomicity()
	fmt.Println("Exiting main()")
}

func testConcurrencyWithoutWait() {
	fmt.Println("Hello concurrency")
	go foo() //launches another go routine
	go bar() //launches another go routine
	fmt.Println("Bye concurrency")
	//program executes without waiting for the child go routines
}

func init() {
	fmt.Println("I am the first function to execute always ")

	//with below line the code will use all the cores to execute the program concurrently
	runtime.GOMAXPROCS(runtime.NumCPU())

}

var wg sync.WaitGroup

func testConcurrencyWithWait() {

	wg.Add(2)

	fmt.Println("Hello concurrency 2")
	go foo() //launches another go routine
	go bar() //launches another go routine
	fmt.Println("Bye concurrency 2")
	wg.Wait()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo : ", i)
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar : ", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done()
}

//global variable
//go run -race main.go to check if the code has race condition or not
var counter int

func testRaceCondition() {
	wg.Add(2)
	fmt.Println("Testing race condition")
	go incrementor("foo")
	go incrementor("bar")
	wg.Wait()
	fmt.Println("Testing race condition done")
}

func incrementor(s string) {
	for i := 0; i < 50; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(5 * time.Millisecond))
		counter = x
		fmt.Println(s, i, " Counter : ", counter)
	}
	wg.Done()
}

var mutex sync.Mutex

func testAvoidRaceConditionWithMutex() {
	wg.Add(2)
	fmt.Println("Testing race condition with mutex")
	go incrementorWithMutex("foo")
	go incrementorWithMutex("bar")
	wg.Wait()
	fmt.Println("Testing race condition done with mutex")
}

func incrementorWithMutex(s string) {
	for i := 0; i < 50; i++ {
		time.Sleep(time.Duration(5 * time.Millisecond))
		mutex.Lock()
		x := counter
		x++
		counter = x
		fmt.Println(s, i, " Counter : ", counter)
		mutex.Unlock()
	}
	wg.Done()
}

var atomicCounter int64

func testWithAtomicity() {
	wg.Add(2)
	fmt.Println("Testing race condition with atomicity")
	go incrementorWithAtomocity("foo")
	go incrementorWithAtomocity("bar")
	wg.Wait()
	fmt.Println("Testing race condition done with atomicity")
}

func incrementorWithAtomocity(s string) {
	for i := 0; i < 500; i++ {
		//time.Sleep(time.Duration(1 * time.Millisecond))
		atomic.AddInt64(&atomicCounter, 1)
		atomic.LoadInt64(&atomicCounter)
		fmt.Println(s, i, " Counter : ", atomicCounter)
	}
	wg.Done()
}
