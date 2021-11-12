package db

import "sync"

type DB struct {
	lock *sync.RWMutex
	Data map[string]Chain // map[guildID]chain
}

func NewChain() *Chain {
	return &Chain{
		lock:  &sync.RWMutex{},
		Links: make(map[string]*Probability),
	}
}

type Chain struct {
	lock  *sync.RWMutex
	Links map[string]*Probability // map[word]possible words
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
