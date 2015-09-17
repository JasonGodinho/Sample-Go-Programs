package main

import (
	"fmt"
	"time"
)

func main() {
	var input string
	// Call the function for thread.after
	z := SleepThread()
	fmt.Println("Seconds thread was asleep:", z)
	// Keep the console open to see the output
	fmt.Scanln(&input)
}

func SleepThread() int {
	// Note the current time
	i := time.Now()
	// Call the thread.after function
	<-time.After(time.Millisecond * 3000)
	//Calculate the time that has lapsed since the program slept
	z := time.Since(i)
	//Return the seconds part of the time
	return int(z.Seconds())
}
