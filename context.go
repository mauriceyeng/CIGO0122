package main

import (
	"context"
	"fmt"
)

//store and pass the info across the different layers of application
//context variable and pass it through the main->router->handler->db function

//the ability to cancellation of the job in between the execution
//(rahul sir's task) consume restful api.. and if not done with it 2minutes then job needs to be cancelled

func main() {
	ctx := context.Background()
	//seed some data in ctx
	ctx = seedContext(ctx)
	readCtx(ctx)
}
func readCtx(ctx context.Context) {
	value := ctx.Value("one")
	fmt.Println(value)
}
func seedContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "one", "111")
	return ctx
}
