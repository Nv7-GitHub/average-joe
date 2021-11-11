package db

import (
	"math/rand"
	"sort"
	"strings"
)

const MaxLoops = 20

func (p *Probability) Predict() string {
	p.lock.Lock()
	defer p.lock.Unlock()
	num := rand.Intn(p.Sum)

	// Sort
	sort.Slice(p.Data, func(i, j int) bool {
		return p.Data[i].Count > p.Data[j].Count
	})

	// Weighted random
	for _, v := range p.Data {
		if num < v.Count {
			return v.Value
		}
		num -= v.Count
	}
	return ""
}

func (c *Chain) Predict() string {
	sentence := &strings.Builder{}

	start := c.Starters.Predict()
	sentence.WriteString(start)

	c.lock.RLock()
	word := c.Chain[start].Predict()
	c.lock.RUnlock()
	loops := 0
	for {
		if word == "EOS" {
			break
		}

		if loops > MaxLoops {
			break
		}
		loops++

		sentence.WriteString(" ")
		sentence.WriteString(word)

		c.lock.RLock()
		word = c.Chain[word].Predict()
		c.lock.RUnlock()
	}

	return sentence.String()
}
