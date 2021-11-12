package joe

import (
	"fmt"

	"github.com/Nv7-Github/average-joe/chain"
	"github.com/bwmarrin/discordgo"
)

// Invite: https://discord.com/oauth2/authorize?client_id=908174272478986311&scope=bot%20applications.commands&permissions=3072

func (b *Bot) MsgCreate(dg *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == dg.State.User.ID {
		return
	}

	b.lock.RLock()
	chn, exists := b.Chains[m.GuildID]
	b.lock.RUnlock()
	if !exists {
		// No chain, create one
		var err error
		chn, err = chain.NewChain(dbPath + m.GuildID + ".txt")
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
