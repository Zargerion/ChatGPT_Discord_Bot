package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"ChatGPT_Discord_Bot/events"
	"ChatGPT_Discord_Bot/commands"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()


	var (
		token string = viper.GetString("DISCORD_TOKEN")

	) 
	if token == ""{
		panic("Missing DISCORD_TOKEN")
	}
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Connection with token. Error creating Discord session,", err)
		return
	}

	//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	commands.InitComandList(discord)

	discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := events.CommandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
    discord.AddHandler(events.Ready)
	discord.AddHandler(events.MessageCreate)

	//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

	err = discord.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	commands.DeleteComandList(discord)

	discord.Close()
	
	log.Println("Gracefully shutting down.")
}