package main

import "syscall/js"

var doc js.Value

func init() {
	doc = js.Global().Get("document")
}

func main() {
	body := doc.Call("querySelector", "body")

	p := doc.Call("createElement", "p")
	p.Set("innerText", "Hello, WASM")

	body.Call("appendChild", p)
}
