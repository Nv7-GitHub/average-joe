package db

import (
	"strings"
)

func (p *Probability) AddWord(word string) {
	p.lock.Lock()
	_, exists := p.Data[word]
	if exists {
		p.Data[word]++
	} else {
		p.Data[word] = 1
	}
	p.Sum++
	p.lock.Unlock()
}

var punctuation = map[byte]empty{
	'.':  {},
	',':  {},
	'?':  {},
	'!':  {},
	'\'': {},
}

func (c *Chain) Add(msg string) {
	msg = strings.ToLower(msg)
	words := strings.Split(msg, " ")

	// Process punctuation - if you had ["hello,", "world!"] you would end up with ["hello", ",", "world", "!"]
	for i, word := range words {
		lastChar := word[len(word)-1]
		_, exists := punctuation[lastChar]
		if exists {
			words[i] = word[:len(word)-1]
			words = insert(words, i+1, string(lastChar))
		}
	}

	// Add words to the chain
	for i, word := range words {
		if i == 0 {
			c.Starters.AddWord(word)
			continue
		}

		_, exists := c.Chain[words[i-1]]
		if !exists {
			c.Chain[words[i-1]] = NewProbability()
		}
		c.Chain[words[i-1]].AddWord(word)

		// Add EOS
		if i == len(words)-1 {
			_, exists := c.Chain[words[i]]
			if !exists {
				c.Chain[words[i]] = NewProbability()
			}
			c.Chain[words[i]].AddWord("EOS")
		}
	}
}
