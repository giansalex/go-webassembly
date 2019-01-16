package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func add(i []js.Value) {
	value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()
	num1, err := strconv.Atoi(value1)
	if err != nil {
		fmt.Println("Invalid Numer: " + value1)
	}
	num2, err := strconv.Atoi(value2)
	if err != nil {
		fmt.Println("Invalid Numer: " + value2)
	}
	result := num1 + num2
	js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", result)
	fmt.Println(result)
}

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(add))
}

func main() {
	c := make(chan struct{}, 0)

	fmt.Println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}