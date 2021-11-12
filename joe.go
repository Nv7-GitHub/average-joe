package joe

import (
	_ "embed"
	"os"
	"strings"
	"sync"

	"github.com/Nv7-Github/average-joe/chain"
	"github.com/bwmarrin/discordgo"
)

const dbPath = "db/averagejoe/"

func NewJoe(token string) (*Bot, error) {
	dg, err := discordgo.New("Bot " + strings.TrimSpace(token))
	if err != nil {
		return nil, err
	}
	err = os.MkdirAll(dbPath, os.ModePerm)
	if err != nil {
		return nil, err
	}

	b := &Bot{
		dg:     dg,
		lock:   &sync.RWMutex{},
		Chains: make(map[string]*chain.Chain),
	}
	b.dg.AddHandler(b.MsgCreate)

	return b, b.dg.Open()
}
