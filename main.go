package main

import "fmt"

func main() {
	body, err := sendGETRequest("https://kamuridesu.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(body)
}
