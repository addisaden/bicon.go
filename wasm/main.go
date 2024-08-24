package main

// GOOS=js GOARCH=wasm go build -o main.wasm
// cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .

import (
	"fmt"
	"syscall/js"

	"github.com/addisaden/bicon.go/base"
)

func main() {
	first_book, err := base.GetBookName(1)
	result := "Hi"
	g1 := base.GetBereanInterlinear1()
	g11 := base.Analyze(g1["1_1_1"])
	fmt.Printf("%v\n", g11)
	if err == nil {
		result = fmt.Sprintf("Hello, WebAssembly: %v", first_book)
	} else {
		result = fmt.Sprintf("Hello, WebAssembly: Error %v", err)
	}
	fmt.Println(result)
	js.Global().Get("console").Call("log", "Hello from go :) Jesus is King!")
}
