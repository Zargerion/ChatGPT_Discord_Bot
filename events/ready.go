package events

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, event *discordgo.Ready) {

	s.UpdateGameStatus(0, "ботоебалю")
	fmt.Println("Fully connected with discord!")
}
