package db

import (
	"strings"
)

func (p *Probability) AddWord(word string) {
	p.lock.Lock()
	_, exists := p.Index[word]
	if exists {
		p.Index[word].Count++
	} else {
		val := &Value{
			Value: word,
			Count: 1,
		}
		p.Index[word] = val
		p.Data = append(p.Data, val)
	}
	p.Sum++
	p.lock.Unlock()
}

var punctuation = []byte{'.', ',', '?', '!', '\''}

func (c *Chain) Add(msg string) {
	msg = strings.ToLower(msg)
	// Remove punctuation
	for _, val := range punctuation {
		msg = strings.ReplaceAll(msg, string(val), "")
	}
	words := strings.Split(msg, " ")

	// Add words to the chain
	for i, word := range words {
		if i == 0 {
			c.Starters.AddWord(word)
		} else {
			_, exists := c.Chain[words[i-1]]
			if !exists {
				c.lock.Lock()
				c.Chain[words[i-1]] = NewProbability()
				c.lock.Unlock()
			}
			c.lock.Lock()
			c.Chain[words[i-1]].AddWord(word)
			c.lock.Unlock()
		}

		// Add EOS
		if i == len(words)-1 {
			c.lock.RLock()
			_, exists := c.Chain[word]
			c.lock.RLock()
			if !exists {
				c.lock.Lock()
				c.Chain[words[i]] = NewProbability()
				c.lock.Unlock()
			}
			c.lock.Lock()
			c.Chain[words[i]].AddWord("EOS")
			c.lock.Unlock()
		}
	}
}
