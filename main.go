package main

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, "Hello, World! Requested path is %q", ctx.Path())
	}
	if err := fasthttp.ListenAndServe(":8080", requestHandler); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
}
