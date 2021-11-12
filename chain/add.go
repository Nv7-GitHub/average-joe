package chain

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

func (c *Chain) Add(msg string) error {
	words := strings.Split(simplify(msg), " ")
	if len(words) == 0 {
		return nil
	}

	// Add words to the chain
	for i, word := range words {
		if i == 0 {
			// Starter word
			err := c.AddLink(Starter, word, 1)
			if err != nil {
				return err
			}
		} else {
			err := c.AddLink(words[i-1], word, 1)
			if err != nil {
				return err
			}
		}

		// End of sentence
		if i == len(words)-1 {
			err := c.AddLink(word, Ender, 1)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
