package db

import (
	"strings"
)

const (
	Starter = "SOS" // StartOfSentence
	Ender   = "EOS" // EndOfSentence
)

func (p *Probability) AddWord(word string, count int) {
	p.lock.Lock()
	_, exists := p.Data[word]
	if exists {
		p.Data[word] += count
	} else {
		p.Data[word] = count
	}
	p.Sum++
	p.lock.Unlock()
}

func (c *Chain) AddLink(start, end string, count int) {
	c.lock.Lock()
	_, exists := c.Links[start]
	if !exists {
		c.Links[start] = NewProbability()
	}
	c.Links[start].AddWord(end, count)
	c.lock.Unlock()
}

func (c *Chain) Add(msg string) {
	words := strings.Split(simplify(msg), " ")
	if len(words) == 0 {
		return
	}

	// Add words to the chain
	for i, word := range words {
		if i == 0 {
			// Starter word
			c.AddLink(Starter, word, 1)
		} else {
			c.AddLink(words[i-1], word, 1)
		}

		// End of sentence
		if i == len(words)-1 {
			c.AddLink(word, Ender, 1)
		}
	}
}
