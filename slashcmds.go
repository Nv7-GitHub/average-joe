package joe

import "github.com/bwmarrin/discordgo"

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "predict",
		Type:        discordgo.ChatApplicationCommand,
		Description: "Predict a message based on previous messages!",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "prompt",
				Description: "A prompt for the bot to start predicting with.",
				Required:    false,
			},
		},
	},
}
