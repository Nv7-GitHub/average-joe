package chain

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"sync"
)

func NewChain(file string) (*Chain, error) {
	chain := &Chain{
		lock:  &sync.RWMutex{},
		Links: make(map[string]*Probability),
	}
	err := chain.Load(file)
	return chain, err
}

func (c *Chain) Load(filename string) error {
	c.filename = filename

	_, err := os.Stat(filename)
	// If file doesn't exist, create it
	if os.IsNotExist(err) {
		c.data, err = os.Create(filename)
		return err
	}
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	rd := bufio.NewReader(file)
	var data LinkEntry
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		// Load line
		err = json.Unmarshal(line, &data)
		if err != nil {
			return err
		}

		// Add link
		c.AddLink(data.Start, data.Next, data.Count)
	}

	c.data, err = os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	return err
}

type LinkEntry struct {
	Start string `json:"start"`
	Next  string `json:"next"`
	Count int    `json:"count"`
}

func (c *Chain) AddLink(start, end string, count int) error {
	c.lock.Lock()
	_, exists := c.Links[start]
	if !exists {
		c.Links[start] = NewProbability()
	}
	c.Links[start].AddWord(end, count)
	c.lock.Unlock()

	// Write
	data := LinkEntry{
		Start: start,
		Next:  end,
		Count: count,
	}
	return c.Write(data)
}

func (c *Chain) Write(entry LinkEntry) error {
	dat, err := json.Marshal(entry)
	if err != nil {
		return err
	}
	_, err = c.data.WriteString(string(dat) + "\n")
	return err
}

// Optimize uses the count value to reduce the number of entries
func (c *Chain) Optimize() error {
	c.lock.Lock()
	var err error
	c.data, err = os.Create(c.filename)
	if err != nil {
		return err
	}

	// Write
	for k, v := range c.Links {
		v.lock.RLock()
		for kv, val := range v.Data {
			err = c.Write(LinkEntry{
				Start: k,
				Next:  kv,
				Count: val,
			})
			if err != nil {
				return err
			}
		}
		v.lock.RUnlock()
	}
	c.lock.Unlock()

	return nil
}
