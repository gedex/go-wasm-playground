package main

import (
	"log"
	"net/http"
)

func main() {
	h := http.FileServer(http.Dir("./build"))

	log.Println("serving ./build and listening on http://localhost:8888")
	if err := http.ListenAndServe("localhost:8888", h); err != nil {
		log.Fatal(err)
	}
}
