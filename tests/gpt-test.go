package main

import (
	"ChatGPT_Discord_Bot/gpt"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	api_key := viper.GetString("GPT_API_KEY")
	if api_key == "" {
		panic("Missing GPT_API_KEY.")
	}
	chat_gpt := gpt.NewGPTConnection(api_key)
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		chat_gpt.ToGPT(text)
	}
}