package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	ctx, cancelledCtx := context.WithCancel(ctx)

	pritch := make(chan int)
	go doAnother(ctx, pritch)

	for i:= 0; i < 3 ; i++{
		pritch <-i
	}

	cancelledCtx()

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Do something finished")


}

func doAnother(ctx context.Context, pritch <-chan int) {
	for {
		select {
		case <- ctx.Done():
			if err:= ctx.Err(); err!=nil {
				fmt.Printf("do another error %s\n", err)
			}
			fmt.Println("DoAnother finished")
			return

		case num := <-pritch:
			fmt.Println("received ", num)
		}
	}
}

func main() {
	var ctx context.Context = context.Background()
	ctx = context.WithValue(ctx, "mykey", "myvalue")
	doSomething(ctx)

}
