package main

import (
	"fmt"
	"log"
)

func main() {
	jsonStr, err := Tokenize(`
{
  "name": "iPhone 6s",
  "price": 649.99,
  "isAvailable": true,
  "owner": null
}
	`)

	if err != nil {
		log.Fatal(err)
	}

	printTokens(jsonStr)

	fmt.Println("\n------------------ RESULT ----------------------")

	result, err := Parser(jsonStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
