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
	hasPrompt := false
	dat := i.ApplicationCommandData()
	if len(dat.Options) > 0 {
		prompt = dat.Options[0].StringValue()
		hasPrompt = true
	}
	prediction, ok := chn.Predict(prompt)
	if !ok {
		b.dg.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
			Content: "No response to prompt found!",
		})
		return
	}
	// Put prompt in start
	if hasPrompt {
		prediction = prompt + " " + prediction
	}
	b.dg.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Content: prediction,
	})
}
