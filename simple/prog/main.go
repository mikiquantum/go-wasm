// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"
)

func printWasm(this js.Value, v []js.Value) interface{} {
	fmt.Println("Hello from WASM", v)
	return nil
}

func main() {
	//c := make(chan struct{}, 0)
	a := 12 + 10
	fmt.Println("[WASM] Result of addition is", a)
	//fmt.Println("WASM Go Initialized")

	// register functions
	//js.Global().Set("printWasm", js.FuncOf(printWasm))
	//fmt.Println("Done...")
	//<-c
}
