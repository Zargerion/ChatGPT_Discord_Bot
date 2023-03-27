package events

import "github.com/bwmarrin/discordgo"


func Ready(s *discordgo.Session, event *discordgo.Ready) {

	s.UpdateGameStatus(0, "OpenAI")
}
