package main

import (
	"ChatGPT_Discord_Bot/gpt"
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
	chat_gpt.ToGPT("Hi!")
}