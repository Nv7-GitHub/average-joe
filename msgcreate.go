package joe

import (
	"fmt"

	"github.com/Nv7-Github/average-joe/chain"
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) MsgCreate(dg *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == dg.State.User.ID || m.Author.Bot {
		return
	}

	if m.Content == "!optimize" {
		_, exists := admins[m.Author.ID]
		if exists {
			b.Optimize(m)
			return
		}
	}
	if m.Content == "!reset" {
		_, exists := admins[m.Author.ID]
		if exists {
			b.Reset(m)
			return
		}
	}

	b.lock.RLock()
	chn, exists := b.Chains[m.GuildID]
	b.lock.RUnlock()
	if !exists {
		// No chain, create one
		var err error
		chn, err = chain.NewChain(dbPath + m.GuildID + ".json")
		if err != nil {
			// TODO: Better error handler
			fmt.Println(err)
			return
		}

		b.lock.Lock()
		b.Chains[m.GuildID] = chn
		b.lock.Unlock()
	}

	err := chn.Add(m.Content)
	if err != nil {
		// TODO: Better error handler
		fmt.Println(err)
		return
	}
}
