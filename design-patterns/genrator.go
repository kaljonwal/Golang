// Generator design pattern :=> Function that returns channel
// Generator Pattern is used to generate a sequence of values which is used to produce some output. This pattern is widely used to introduce parallelism into loops.
// This allows the consumer of the data produced by the generator to run in parallel when the generator function is busy computing the next value.

package main

import (
	"fmt"
	"time"
)

func funcWhichReturnsChannel(n int) chan int {
	ch := make(chan int)
	go func() {
		time.Sleep(1000 * time.Millisecond)
		ch <- n
	}()

	return ch
}

func main() {
	ch := funcWhichReturnsChannel(8)

	fmt.Println("waiting")
	fmt.Println(<-ch)
}
