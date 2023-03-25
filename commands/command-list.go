package commands

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	defaultPermission bool = true
	commands = []*discordgo.ApplicationCommand {
		{
			Name: "ping",
			Description: "Checks working.",
			DefaultPermission: &defaultPermission,
		},
		{
			Name: "text",
			Description: "To write and get text.",
			DefaultPermission: &defaultPermission,
			Options: []*discordgo.ApplicationCommandOption{

				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "message",
					Description: "Write",
					Required:    true,
				},
			},
		},	
		{
			Name: "image",
			Description: "To write text but get text with image. *Can be unavalable.",
			DefaultPermission: &defaultPermission,
		},
    }
	registeredCommands []*discordgo.ApplicationCommand
)

func InitComandList(s *discordgo.Session) {
	log.Println("Adding commands...")
	registeredCommands = make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate("1088457644706103296", "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
}

func DeleteComandList(s *discordgo.Session)  {
	log.Println("Deleting commands...")
	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
		
	}
}