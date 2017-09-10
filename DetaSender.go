package main

import (
	"fmt"
	"strconv"
	"github.com/valyala/fasthttp"
)

func fastHTTPHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/plain")

	path := string(ctx.Path())
	size, err := strconv.Atoi(path[1:])
	if err != nil {
		ctx.Error("Unsupported path", fasthttp.StatusNotFound)
	} else {
		var res string
		for i := 0; i < size; i++ {
			res += "0"
		}
		fmt.Fprintf(ctx, "%s", res)
	}
}

func main() {
	port := "8080"
	fasthttp.ListenAndServe(":"+port, fastHTTPHandler)
}
