package gpt

import (
	"ChatGPT_Discord_Bot/gpt/interfaceGPT"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

var (
	api_key string
	clearSignal bool = false
	chat_gpt interfaceGPT.IGPT
	plus_line string = "++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++"

)
	
func init() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	api_key = viper.GetString("GPT_API_KEY")
	if api_key == "" {
		viper.AutomaticEnv()
		err := viper.BindEnv("GPT_API_KEY", "GPT_API_KEY")
		if err != nil {
			fmt.Println("Missing api key for discord.")
			return
		}
		api_key = viper.GetString("GPT_API_KEY")	
	}
	chat_gpt = interfaceGPT.NewGPT(api_key)
}


func Chat20Requests(msg *string, s *discordgo.Session, i *discordgo.InteractionCreate, NumOfCallText int8) {

	if NumOfCallText == 19 {
		clearSignal = true
	}

	ans, err := chat_gpt.ToGPTWithHistoryChat(*msg, clearSignal)

	if err != nil {
		fmt.Println("Error: Cannot send/get message.")
		ans = "Error: Cannot send/get message."
	} else {
		fmt.Println((plus_line + "\n\n" + ans + "\n\n" + plus_line + "\n@" + i.Member.User.Username))
		s.ChannelMessageSend(i.ChannelID, (plus_line + "\n\n" + ans + "\n\n" + plus_line + "\n<@" + i.Member.User.ID + ">")) 
	}
}



func SendToTranslateWithGrammarCorrection(msg *string, s *discordgo.Session, i *discordgo.InteractionCreate) {

	ans, err := chat_gpt.ToGPTWithoutHistoryTranslate(*msg)

	if err != nil {
		fmt.Println("Error: Cannot send/get message for translate.")
		ans = "Error: Cannot send/get message for translate."
	} else {
		fmt.Println(("Translated:\n" + plus_line + "\n\n" + ans + "\n\n" + plus_line + "\n@" + i.Member.User.Username))
		*msg = ans
	}

	ans, err = chat_gpt.ToGPTWithoutHistoryGrammarCorrect(*msg)

	if err != nil {
		fmt.Println("Error: Cannot send/get message for grammar correction.")
		ans = "Error: Cannot send/get message for grammar correction."
	} else {
		fmt.Println(("Grammar corrected:\n" + plus_line + ans + "\n\n" + plus_line + "\n@" + i.Member.User.Username))
		s.ChannelMessageSend(i.ChannelID, (plus_line + ans + "\n\n" + plus_line + "\n<@" + i.Member.User.ID + ">")) 
	}
}

func DeletingChatMessages() {
	msg := new(string)
	*msg = "Bye!"
	chat_gpt.ToGPTWithHistoryChat(*msg, true)
}