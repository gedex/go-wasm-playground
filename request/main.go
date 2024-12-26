package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	request("GET", "https://api.github.com/users/gedex")
	request("GET", "https://api.github.com/not_found")
}

func request(method, url string) {
	log.Printf("%s %s", method, url)

	c := http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("resp %s", string(b))
}
