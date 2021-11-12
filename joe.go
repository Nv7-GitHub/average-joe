package joe

import (
	_ "embed"
	"os"
	"strings"
	"sync"

	"github.com/Nv7-Github/Nv7Haven/dgutil"
	"github.com/Nv7-Github/average-joe/chain"
	"github.com/bwmarrin/discordgo"
)

const dbPath = "data/averagejoe/"
const clientID = "908174272478986311"

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

	// Load data
	files, err := os.ReadDir(dbPath)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		gld := strings.Split(file.Name(), ".")[0]
		chain, err := chain.NewChain(dbPath + file.Name())
		if err != nil {
			return nil, err
		}
		b.Chains[gld] = chain
	}

	// Handlers
	b.dg.AddHandler(b.MsgCreate)
	b.dg.AddHandler(b.predict)
	dgutil.UpdateBotCommands(b.dg, clientID, "", commands)

	return b, b.dg.Open()
}
