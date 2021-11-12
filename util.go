package joe

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

type empty struct{}

var admins = map[string]empty{
	"567132457820749842": {},
}

func (b *Bot) Optimize(m *discordgo.MessageCreate) {
	msg, err := b.dg.ChannelMessageSendReply(m.ChannelID, "Optimizing...", m.Reference())
	if err != nil {
		fmt.Println(err)
		return
	}
	start := time.Now()

	b.lock.RLock()
	for _, chain := range b.Chains {
		err := chain.Optimize()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	b.lock.RUnlock()

	b.dg.ChannelMessageEdit(m.ChannelID, msg.ID, "Optimized in "+time.Since(start).String()+".")
}

func (b *Bot) Reset(m *discordgo.MessageCreate) {
	b.lock.RLock()
	err := b.Chains[m.GuildID].Reset()
	b.lock.RUnlock()
	if err != nil {
		b.dg.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	b.dg.ChannelMessageSend(m.ChannelID, "Reset server!")
}
