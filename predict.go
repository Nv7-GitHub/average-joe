package joe

import (
	"github.com/Nv7-Github/average-joe/chain"
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) predict(s *discordgo.Session, i *discordgo.InteractionCreate) {
	b.lock.RLock()
	chn, exists := b.Chains[i.GuildID]
	b.lock.RUnlock()
	if !exists {
		b.dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags:   1 << 6,
				Content: "No messages have been sent yet!",
			},
		})
	}

	// Defer to get loading indicator
	b.dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
	})

	prompt := chain.Starter
	dat := i.ApplicationCommandData()
	if len(dat.Options) > 0 {
		prompt = dat.Options[0].StringValue()
	}
	prediction, ok := chn.Predict(prompt)
	if !ok {
		b.dg.FollowupMessageCreate(clientID, i.Interaction, true, &discordgo.WebhookParams{
			Content: "No response to prompt found!",
		})
		return
	}
	b.dg.FollowupMessageCreate(clientID, i.Interaction, true, &discordgo.WebhookParams{
		Content: prompt + " " + prediction,
	})
}
