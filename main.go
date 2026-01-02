package main

import "log"

func main() {
	jsonStr, err := Tokenize(`
{
  "name": "iPhone 6s",
  "price": 649.99,
  "isAvailable": true
}
	`)

	if err != nil {
		log.Fatal(err)
	}

	printTokens(jsonStr)
}
