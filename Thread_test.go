package Thread_test

import (
	"testing"
	"time"
)

func TestPrintNumbers(t *testing.T) {
	for _, value := range expected {
		//Test the function with the input and output values array
		k := SleepThread(value.input)
		// If the output doesn't match, show that the test failed
		if k != value.output {
			t.Errorf("Thread slept for %d seconds instead of  %d", value.input, value.output)
		}
	}
}
func SleepThread(a int) int {
	// Note the current time
	i := time.Now()
	// Call the thread.after function
	<-time.After(time.Second * time.Duration(a))
	//Calculate the time that has lapsed since the program slept
	z := time.Since(i)
	//Return the seconds part of the time
	return int(z.Seconds())
}

//Struct for input and output values
type abc struct {
	input  int
	output int
}

//Expected input and output values for testing
var expected = []abc{
	{1, 1},
	{2, 2},
	{3, 3},
}
