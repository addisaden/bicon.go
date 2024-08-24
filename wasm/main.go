package main

// GOOS=js GOARCH=wasm go build -o main.wasm
// cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/addisaden/bicon.go/base"
)

func jsGetBookName(this js.Value, args []js.Value) any {
	if len(args) >= 1 {
		name, err := base.GetBookName(args[0].Int())
		if err == nil {
			return name
		}
	}
	return nil
}

func jsGetBookIndex(this js.Value, args []js.Value) any {
	if len(args) >= 1 {
		name_index, err := base.GetBookIndex(args[0].String())
		if err == nil {
			return name_index
		}
	}
	return nil
}

func jsGetTexts(this js.Value, args []js.Value) any {
	if len(args) >= 4 {
		book := uint8(args[0].Int())
		chapter := args[1].Int()
		from_verse := args[2].Int()
		to_verse := args[3].Int()

		book_data := base.GetBooks(book)
		result := map[string]base.VerseTextType{}
		for i := from_verse; i <= to_verse; i++ {
			this_index := fmt.Sprintf("%d_%d_%d", book, chapter, i)
			result[this_index] = base.Analyze(book_data[this_index])
		}
		jsonBytes, _ := json.Marshal(result)
		jsValue := js.Global().Get("JSON").Call("parse", string(jsonBytes))
		return jsValue
	}
	return nil
}

func main() {
	first_book, err := base.GetBookName(1)
	result := "Hi"
	g1 := base.GetBereanInterlinear()
	g11 := base.Analyze(g1["1_1_1"])
	fmt.Printf("%v\n", g11)
	if err == nil {
		result = fmt.Sprintf("Hello, WebAssembly: %v", first_book)
	} else {
		result = fmt.Sprintf("Hello, WebAssembly: Error %v", err)
	}
	fmt.Println(result)
	js.Global().Get("console").Call("log", "Hello from go :) Jesus is King!")

	js.Global().Set("getBookName", js.FuncOf(jsGetBookName))
	js.Global().Set("getBookIndex", js.FuncOf(jsGetBookIndex))
	js.Global().Set("getTexts", js.FuncOf(jsGetTexts))
	<-make(chan bool)
}
