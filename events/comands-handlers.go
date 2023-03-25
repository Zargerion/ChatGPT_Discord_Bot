package events

import (
	"fmt"
	"log"

	"ChatGPT_Discord_Bot/utils"

	"github.com/bwmarrin/discordgo"
)

var (
	GottenInfo string
	CommandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		"ping": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			id_message := i.Interaction.ID
			timestamp, err := discordgo.SnowflakeTimestamp(id_message)
			if err != nil {
				log.Panicf("Error SnowflakeTimestamp(id_message) %s", err)
			}
			timestamp = timestamp.UTC()
			delta, err := utils.TimeDelta(&timestamp)
			if err != nil {
				fmt.Println(err)
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Got it! Your ping is " + fmt.Sprint(delta) + " seconds.",
				},
			})
		},
		"text": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

			options := i.ApplicationCommandData().Options

			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			margs := make([]interface{}, 0, len(options))
			msgformat := ""

			if option, ok := optionMap["message"]; ok {

				margs = append(margs, option.StringValue())
				msgformat += "%s"
			}

			_, err := fmt.Println(fmt.Sprintf(
				msgformat,
				margs...,
			))
			if err != nil {
				fmt.Println("Error: fmt.Println(msgformat).")
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Got it! You sent: " + fmt.Sprintf(
						msgformat,
						margs...,
					),
				},
			})
		},
		"image": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Got it, but now it's unavalable. We are needs GPT Plus...",
				},
			})
		},
    }
)