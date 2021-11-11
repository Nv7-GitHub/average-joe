package db

import "sync"

type empty struct{}

type DB struct {
	lock *sync.RWMutex
	Data map[string]Chain // map[guildID]chain
}

func NewChain() *Chain {
	return &Chain{
		lock:     &sync.RWMutex{},
		Chain:    make(map[string]*Probability),
		Starters: NewProbability(),
	}
}

type Chain struct {
	lock     *sync.RWMutex
	Chain    map[string]*Probability // map[word]possible words
	Starters *Probability
}

func NewProbability() *Probability {
	return &Probability{
		lock: &sync.RWMutex{},
		Data: make(map[string]int),
	}
}

type Probability struct {
	lock *sync.RWMutex
	Data map[string]int // map[word]times said
	Sum  int
}
