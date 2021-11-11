package db

import "sync"

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
		lock:  &sync.RWMutex{},
		Index: make(map[string]*Value),
		Data:  make([]*Value, 0),
	}
}

type Probability struct {
	lock  *sync.RWMutex
	Data  []*Value
	Index map[string]*Value
	Sum   int
}

type Value struct {
	Value string
	Count int
}
