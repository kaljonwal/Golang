package main

import (
	"fmt"
	"sync"
)

var counter = 0

func main() {
	var once sync.Once

	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println("triggered")
			once.Do(performoperationOnce)
		}()
	}

}

func performoperationOnce() {
	counter++
	fmt.Println(counter)
}
