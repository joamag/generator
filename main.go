package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Printer(channel chan int, delay time.Duration, count int) {
	// iterates over the requested count of
	// generates numbers to send them to the
	// target channel for processing
	for index := 0; index < count; index++ {
		channel <- rand.Int()
		time.Sleep(delay)
	}

	// closes the channel at the end of the number
	// generation as defined in specification
	close(channel)
}

func main() {
	// prints a message about the starting of the
	// process (required for debuggin)
	fmt.Println("Starting up")

	// creates the channel that is going to be
	// used in the communication between both of
	// the routine structures
	channel := make(chan int)

	// starts the printer co-routine that is
	// going to be used
	go Printer(channel, 1*time.Second, 10)

	// iterates continuously while the valid
	// flag is set (to receive all of the data)
	for valid := true; valid; {
		// waits for the communication from the other
		// side of the co-routine that is going to send
		// some information (as requested)
		select {
		case value, ok := <-channel:
			if ok {
				fmt.Println("Received := " + strconv.Itoa(value))
			} else {
				valid = false
			}
			break
		}
	}

	// prints a message about the finishing of the
	// execution of the producer consumer
	fmt.Println("Finished execution")
}