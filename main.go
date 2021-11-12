package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Nv7-Github/average-joe/chain"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	chain, err := chain.NewChain("data.txt")
	if err != nil {
		panic(err)
	}

	/*
		sentences := []string{
			//"Hello, World!",
			"hi my name is joe",
			"hi i like food",
			"hi my name is john",
			"hello",
			//"lol 😳",
			//"lol",
		}
		for _, sentence := range sentences {
			err = chain.Add(sentence)
			if err != nil {
				panic(err)
			}
		}
	*/

	err = chain.Optimize()
	if err != nil {
		panic(err)
	}

	fmt.Println(chain.Predict())
}
