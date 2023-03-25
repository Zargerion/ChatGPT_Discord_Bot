package gpt

import (
	"context"
	"fmt"
	"log"

	"github.com/PullRequestInc/go-gpt3"
)

type IGPT interface {
	ToGPT(text string) (string, error)
}

type igpt struct {
	gptClint *gpt3.Client
	ctx      *context.Context
}

func NewGPTConnection(apiKey string) IGPT {

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
		log.Fatalln(err)
	} else {
		fmt.Print("Gpt is running.")
	}

	return gpt
}

var messages = make([]gpt3.ChatCompletionRequestMessage, 0)

func (gpt *igpt) ToGPT(text string) (string, error) {

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
		MaxTokens:    		180,
		TopP:         		1,
		PresencePenalty:    0.0,
		FrequencyPenalty:   0.6,
		N:            		1,
		Stop:         		[]string{" Human:", " AI:"},
		User:               "user",
		Stream:             false,
	}

	ans := ""

	err := c.ChatCompletionStream(ctx, request, func(resp *gpt3.ChatCompletionStreamResponse){
		fmt.Print(resp.Choices[0].Delta.Content)
		ans = resp.Choices[0].Delta.Content
	})
	if err != nil {
		log.Fatalln(err)
	}

	messages = append(messages, gpt3.ChatCompletionRequestMessage {
	
		Role:    "user",
		Content: ans,	
	})

	return ans, err
}
