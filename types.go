package joe

import (
	"sync"

	"github.com/Nv7-Github/average-joe/chain"
	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	lock *sync.RWMutex

	dg     *discordgo.Session
	Chains map[string]*chain.Chain // map[guild]chain
}

func (b *Bot) Close() {
	b.dg.Close()
	for _, chain := range b.Chains {
		chain.Close()
	}
}
