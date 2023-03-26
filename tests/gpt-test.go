package main

import (
	"ChatGPT_Discord_Bot/gpt"
	"bufio"
	"fmt"
	"log"
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
	chat_gpt := gpt.NewGPTConnection(api_key)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\n")
	fmt.Print("-> ")
	fmt.Print("\n")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	//chat_gpt.ToGPT(text)



	outCh := make(chan string)
	errCh := make(chan error)
	go func() {
		_, err := chat_gpt.ToGPT(text)
		outCh <- "sdf"
		errCh <- err
		close(outCh)
		close(errCh)
	}()
	err := <-errCh
	if err != nil{
		log.Panic(err)
		
	} else {
		out := <-outCh
		fmt.Println(out)
		return
	}
		
	
	//for {
	//	fmt.Print("\n")
	//	fmt.Print("-> ")
	//	text, _ := reader.ReadString('\n')
	//	text = strings.Replace(text, "\n", "", -1)
	//	chat_gpt.ToGPT(text)
	//}
}