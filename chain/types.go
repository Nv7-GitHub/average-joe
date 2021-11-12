package chain

import (
	"os"
	"sync"
)

type Chain struct {
	lock  *sync.RWMutex
	Links map[string]*Probability // map[word]possible words

	data     *os.File
	filename string
}

func NewProbability() *Probability {
	return &Probability{
		lock: &sync.RWMutex{},
		Data: make(map[string]int),
	}
}

type Probability struct {
	lock *sync.RWMutex
	Data map[string]int
	Sum  int
}
