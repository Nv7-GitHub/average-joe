package db

import (
	"math/rand"
	"strings"
)

const MaxLoops = 20

func (p *Probability) Predict() string {
	p.lock.RLock()
	defer p.lock.RUnlock()
	num := rand.Intn(p.Sum)

	// Weighted random
	for k, v := range p.Data {
		if v <= num {
			return k
		}
	}

	return ""
}

func (c *Chain) Predict() string {
	sentence := &strings.Builder{}

	start := c.Starters.Predict()
	word := c.Chain[start].Predict()
	loops := 0
	for {
		if loops > MaxLoops {
			break
		}
		loops++

		sentence.WriteString(word)
		word = c.Chain[word].Predict()
		if word == "EOS" {
			break
		}
	}

	return sentence.String()
}
