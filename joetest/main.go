package main

import (
	_ "embed"

	joe "github.com/Nv7-Github/average-joe"
)

//go:embed token.txt
var token string

func main() {
	joe, err := joe.NewJoe(token)
	if err != nil {
		panic(err)
	}
	defer joe.Close()
}
