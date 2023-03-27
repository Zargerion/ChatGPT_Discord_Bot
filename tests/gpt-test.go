package main

import (
	"ChatGPT_Discord_Bot/gpt/interfaceGPT"
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
		viper.AutomaticEnv()
		err := viper.BindEnv("GPT_API_KEY", "GPT_API_KEY")
		if err != nil {
			fmt.Println("Missing api key for discord.")
			return
		}
		api_key = viper.GetString("GPT_API_KEY")	
	}
	chat_gpt := interfaceGPT.NewGPT(api_key)
	reader := bufio.NewReader(os.Stdin)
	signal := false
	var count int8 = 0
	for {
		fmt.Print("\n")
		fmt.Print("-> ")
		fmt.Print("\n")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		helpCh := make(chan string)
		count++
		if count == 2 {
			signal = true
		}
		go func() {
			str, err := chat_gpt.ToGPT(text, signal)
			if err != nil{
				fmt.Println("Out fail.")
				return
			} else {
				helpCh <- str
			}
		}() 
		fmt.Print(<-helpCh)
	}
}