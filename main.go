package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Nv7-Github/average-joe/db"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	sentences := []string{
		//"Hello, World!",
		"hi my name is joe",
		"hi i like food",
		"hi my name is john",
		"hello",
		//"lol 😳",
		//"lol",
	}
	chain := db.NewChain()
	for _, sentence := range sentences {
		chain.Add(sentence)
	}

	fmt.Println(chain.Predict())
}
