package chain

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

func (c *Chain) Predict(prompt string) (string, bool) {
	if prompt != Starter {
		split := strings.Split(simplify(prompt), " ")
		prompt = split[len(split)-1]

		c.lock.RLock()
		_, exists := c.Links[prompt]
		c.lock.RUnlock()
		if !exists {
			return "", false
		}
	}
	sentence := &strings.Builder{}

	start := c.Links[prompt].Predict()
	if start == Ender {
		return "", true
	}
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

	return sentence.String(), true
}
