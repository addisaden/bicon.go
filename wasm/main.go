package main

// GOOS=js GOARCH=wasm go build -o main.wasm
// cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WebAssembly")
	js.Global().Get("console").Call("log", "Hello from go :) Jesus is King!")
}
