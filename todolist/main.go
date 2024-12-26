package main

import (
	"syscall/js"
)

var (
	todos     []string
	todosList js.Value
)

func init() {
	todos = make([]string, 0)
	js.Global().Set("addTodo", js.FuncOf(addTodo))
	todosList = js.Global().Get("document").Call("getElementById", "todos")
}

func main() {
	// Prevent exit
	select {}
}

func addTodo(_ js.Value, args []js.Value) any {
	todo := args[0].String()
	todos = append(todos, todo)
	renderTodo(todo)

	return nil
}

func renderTodo(todo string) {
	li := js.Global().Get("document").Call("createElement", "li")
	li.Set("innerText", todo)
	todosList.Call("appendChild", li)
}
