package events

import (
	"ChatGPT_Discord_Bot/gpt"

	"github.com/bwmarrin/discordgo"
)

var (
	ComponentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		"ClearChatHistoryButton": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			if i.Type == discordgo.InteractionMessageComponent {

				NumOfCallText = 0

				gpt.DeletingChatMessages()
	
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "History cleared.",
					},
				})
			}
		},
	}	
)

