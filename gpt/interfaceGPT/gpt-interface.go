package interfaceGPT

import (
	"context"
	"fmt"
	"log"

	"github.com/PullRequestInc/go-gpt3"
)

type IGPT interface {
	ToGPTWithHistoryChat(text string, clearSignal bool) (string, error)
	ToGPTWithoutHistoryTranslate(text string) (string, error)
	ToGPTWithoutHistoryGrammarCorrect(text string) (string, error)
}

type igpt struct {
	gptClint *gpt3.Client
	ctx      *context.Context
}

func NewGPT(apiKey string) IGPT {

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	gpt := &igpt{
		gptClint: &client,
		ctx:      &ctx,
	}

	_, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt: []string{
			"request",
		},
		MaxTokens: gpt3.IntPtr(5),
		Stop:      []string{"."},
		Echo:      true,
	})
	if err != nil {
		log.Fatalln("Cannot make client OpenAI. ", err)
	} else {
		fmt.Print("Gpt client session is started. \n")
	}

	return gpt
}
 
var messages = make([]gpt3.ChatCompletionRequestMessage, 0)

func (gpt *igpt) ToGPTWithHistoryChat(text string, clearSignal bool) (string, error) {

	c := *gpt.gptClint
	ctx := *gpt.ctx

	messages = append(messages, gpt3.ChatCompletionRequestMessage {
	
		Role:    "user",
		Content: text,	
	})

	request := gpt3.ChatCompletionRequest{
		Model:        		"gpt-3.5-turbo",
		Messages:     		messages,
		Temperature:  		gpt3.Float32Ptr(0.9),
		MaxTokens:    		300,
		TopP:         		1,
		PresencePenalty:    0.0,
		FrequencyPenalty:   0.6,
		N:            		1,
		Stop:         		[]string{" Human:", " AI:"},
		User:               "user",
		Stream:             false,
	}

	out := new(string)

	err := c.ChatCompletionStream(ctx, request, func(resp *gpt3.ChatCompletionStreamResponse) {
		*out += resp.Choices[0].Delta.Content
	})
	if err != nil {
		log.Panicln("Error to request.", err)
		return "Error to request.", err
	}
	ans := *out

	messages = append(messages, gpt3.ChatCompletionRequestMessage {
	
		Role:    "user",
		Content: ans,	
	})

	if clearSignal {
		messages = make([]gpt3.ChatCompletionRequestMessage, 0)
		fmt.Println("Chat messages is cleared.")
	}

	return ans, err
}

func (gpt *igpt) ToGPTWithoutHistoryTranslate(text string) (string, error) {

	c := *gpt.gptClint
	ctx := *gpt.ctx

	message := []string {
		"Translate this into English:\n" + text,	
	}

	request := gpt3.CompletionRequest{
		Prompt:     		message,
		Temperature:  		gpt3.Float32Ptr(0.3),
		MaxTokens:    		gpt3.IntPtr(300),
		TopP:         		gpt3.Float32Ptr(1),
		PresencePenalty:    0.0,
		FrequencyPenalty:   0.0,
		Stream:             false,
	}

	out, err := c.CompletionWithEngine(ctx, "text-davinci-003", request)
	if err != nil {
		log.Panicln("Error to request.", err)
		return "Error to request.", err
	}
	ans := out.Choices[0].Text

	return ans, err
}

func (gpt *igpt) ToGPTWithoutHistoryGrammarCorrect(text string) (string, error) {

	c := *gpt.gptClint
	ctx := *gpt.ctx

	message := []string {
		"Correct this to standard English:\n" + text,	
	}

	request := gpt3.CompletionRequest{
		Prompt:     		message,
		Temperature:  		gpt3.Float32Ptr(0),
		MaxTokens:    		gpt3.IntPtr(300),
		TopP:         		gpt3.Float32Ptr(1),
		PresencePenalty:    0.0,
		FrequencyPenalty:   0.0,
		Stream:             false,
	}

	out, err := c.CompletionWithEngine(ctx, "text-davinci-003", request)
	if err != nil {
		log.Panicln("Error to request.", err)
		return "Error to request.", err
	}
	ans := out.Choices[0].Text

	return ans, err
}
