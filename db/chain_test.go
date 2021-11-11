package db

import (
	"fmt"
	"testing"
)

func TestChain(t *testing.T) {
	sentences := []string{
		"Hello, World!",
		//"hi my name is joe",
		//"lol 😳",
	}
	chain := NewChain()
	for _, sentence := range sentences {
		chain.Add(sentence)
	}
	fmt.Println(chain.Predict())
}
