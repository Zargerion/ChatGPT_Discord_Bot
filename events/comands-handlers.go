package events

import (
	"fmt"
	"log"

	"ChatGPT_Discord_Bot/gpt"
	"ChatGPT_Discord_Bot/utils"

	"github.com/bwmarrin/discordgo"
)

var (
	GottenInfo      string
	NumOfCallText   int8   = 0
	storyCleared    string = ""
	CommandHandlers        = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
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
		"chat": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

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

			message := fmt.Sprintf(
				msgformat,
				margs...,
			)

			NumOfCallText++
			if NumOfCallText == 19 {
				storyCleared = "Your message history has reached the limit. After this answer it will be cleared. "
			}
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: `Got it! You sent: "` + message + `" ` + storyCleared + "Wait for answer...",
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{
									Label:    "Сlear history",
									Style:    discordgo.SecondaryButton,
									Disabled: false,
									CustomID: "ClearChatHistoryButton",
								},
							},
						},
					},
				},
			})

			go gpt.Chat20Requests(&message, s, i, NumOfCallText)

			if NumOfCallText == 19 {
				NumOfCallText = 0
			}
		},
		"translate": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

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

			message := fmt.Sprintf(
				msgformat,
				margs...,
			)

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: `Got it! You sent: "` + message + `" ` + "Wait for translation and grammar correction...",
				},
			})

			go gpt.SendToTranslateWithGrammarCorrection(&message, s, i)

		},
		"image": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Got it, but now it's unavalable. We needs GPT Plus..." + " <@" + i.Member.User.ID + ">",
				},
			})
		},
		"dan": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

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

			message := fmt.Sprintf(
				msgformat,
				margs...,
			)

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: `Got it! You sent: "` + message + `" ` + "Wait for DEN...",
				},
			})

			go gpt.SendToDAN(&message, s, i)

		},
	}
)
