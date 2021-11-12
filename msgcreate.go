package joe

import (
	"fmt"

	"github.com/Nv7-Github/average-joe/chain"
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) MsgCreate(dg *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == dg.State.User.ID {
		return
	}

	b.lock.RLock()
	chn, exists := b.Chains[m.GuildID]
	b.lock.RUnlock()
	if !exists {
		// No chain, create one
		chn, err := chain.NewChain(dbPath + m.GuildID + ".txt")

		b.lock.Lock()
		b.Chains[m.GuildID] = chn
		b.lock.Unlock()

		if err != nil {
			// TODO: Better error handler
			fmt.Println(err)
			return
		}
	}

	err := chn.Add(m.Content)
	if err != nil {
		// TODO: Better error handler
		fmt.Println(err)
		return
	}
}
