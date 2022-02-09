package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	go doSomething(ctx)
	select {
	case v := <-ctx.Done():
		fmt.Println("timeline exceeded of 2seconds", v)
	}
	time.Sleep(time.Second * 3)

}

// with error statement
func doSomething(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			err := ctx.Err()
			fmt.Println(err)
			return
		default:
			fmt.Println("Doing something bakwaas")
		}
	}

}
