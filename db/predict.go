package db

import (
	"math/rand"
	"strings"
)

const MaxLoops = 20

func (p *Probability) Predict() string {
	p.lock.Lock()
	defer p.lock.Unlock()
	num := rand.Intn(p.Sum)

	// Weighted random
	for k, v := range p.Data {
		if num < v {
			return k
		}
		num -= v
	}
	return ""
}

func (c *Chain) Predict() string {
	sentence := &strings.Builder{}

	start := c.Links[Starter].Predict()
	sentence.WriteString(start)

	c.lock.RLock()
	word := c.Links[start].Predict()
	c.lock.RUnlock()
	loops := 0
	for {
		if word == Ender {
			break
		}

		if loops > MaxLoops {
			break
		}
		loops++

		sentence.WriteString(" ")
		sentence.WriteString(word)

		c.lock.RLock()
		word = c.Links[word].Predict()
		c.lock.RUnlock()
	}

	return sentence.String()
}
