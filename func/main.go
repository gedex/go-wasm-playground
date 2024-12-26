package main

import (
	"syscall/js"
)

func main() {
	js.Global().Set("add", js.FuncOf(add))

	// keep running
	select {}
}

func add(_ js.Value, args []js.Value) any {
	a := args[0].Int()
	b := args[1].Int()

	return a + b
}
