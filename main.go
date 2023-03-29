package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"ChatGPT_Discord_Bot/commands"
	"ChatGPT_Discord_Bot/events"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	var token string = viper.GetString("DISCORD_TOKEN")

	if token == ""{
		viper.AutomaticEnv()
		err := viper.BindEnv("GDISCORD_TOKEN", "DISCORD_TOKEN")
		if err != nil {
			fmt.Println("Missing api key for discord.")
			return
		}
		token = viper.GetString("DISCORD_TOKEN")	
	}
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Connection with token. Error creating Discord session,", err)
		return
	} else {
		fmt.Println("Discord client session is started.") 
	}

	//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	discord.AddHandler(events.Ready)
	discord.AddHandler(events.MessageCreate)
	commands.InitComandList(discord)
	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:

			if h, ok := events.CommandHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}

		case discordgo.InteractionMessageComponent:
			
			if h, ok := events.ComponentsHandlers[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			}
		}
	})

	//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	err = discord.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
		return
	}

	fmt.Print("\nPress CTRL-C to exit.\n\n")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	commands.DeleteComandList(discord)

	discord.Close()
	
	fmt.Println("Gracefully shutting down.")
}