package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func example1() {
	// Step 1 - This example will give a deadlock error because we are not using goroutine to send and receive data from channel
	// we are not using buffered channel
	// Deadlock is happened due to line no. 12 n := <-dataChan because it will be waiting for data to be received from channel
	dataChan := make(chan int) // In place of int, you can use any type like string, float64, etc. or any struct

	dataChan <- 10 // Send data to channel
	n := <-dataChan
	fmt.Println("Data received from channel:", n)
}

func example2() {
	dataChan := make(chan int, 1) // 1 is the buffer size of channel
	dataChan <- 10                // Send data to channel
	n := <-dataChan
	fmt.Println("Data received from channel:", n)

}

func example3() {
	dataChan := make(chan int) // Unbuffered channel
	go func() {
		dataChan <- 10 // Send data to channel
	}()
	n := <-dataChan //  Here go routine and main routine are running parallelly
	fmt.Println("Data received from channel:", n)
}

func example4() {
	dataChan := make(chan int)
	go func() { // This will be running as background thread
		for i := 0; i < 10; i++ {
			dataChan <- i
		}
		close(dataChan) // This will close the channel and avoid deadlock error in main thread for loop
	}()
	for n := range dataChan { // This will be running as main thread
		fmt.Println("Data received from channel:", n)
	}
	// This will be end up in deadlock because we are running for loop in main thread and it will be waiting for data to be received from channel
}

func doSomeReleasticWork() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(100)
}

func example5() {
	dataChan := make(chan int)
	go func() {
		for i := 0; i < 10; i++ { // This is calling doSomeReleasticWork() function one by one and wait for 1 second to get result
			result := doSomeReleasticWork()
			dataChan <- result
		}
		close(dataChan)
	}()
	for n := range dataChan {
		fmt.Println("Data received from channel:", n)
	}

}

func example6() {
	dataChan := make(chan int)

	go func() {
		wg := sync.WaitGroup{} // This is used to wait for all goroutines to finish. Create a wait group

		for i := 0; i < 10; i++ {
			wg.Add(1) // Add 1 to wait group. Basically it will add 1 to wait group counter
			go func() {
				defer wg.Done() // This will be called when goroutine is finished. Basically it will subtract 1 from wait group counter
				result := doSomeReleasticWork()
				dataChan <- result
			}()

		}
		wg.Wait() // This will wait for all goroutines to finish
		close(dataChan)
	}()
	for n := range dataChan {
		fmt.Println("FAST Data received from channel:", n)
	}

}

func main() {

	// Channel is like pipe , it is used to send and receive data between goroutines
	// Channel is type of data structure
	// Channel is used to communicate between goroutines and multithreading programming

	// Example 1
	//example1()

	// Example 2 - Passing buffer size to channel
	//example2()

	// Example 3 - Using goroutine to send and receive data from channel
	//example3()

	// Example 4 - Running for loop to push and pop data from channel
	// example4()

	// Example 5 Realistic Example - Using channel to send and receive data from channel [Slow Version]
	//example5()

	// Example 6 Realistic Example - Using channel to send and receive data from channel [Fast Version]
	example6()

}


