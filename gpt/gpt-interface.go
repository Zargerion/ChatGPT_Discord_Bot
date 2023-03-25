package gpt

import (
	"context"
	"fmt"
	"log"

	"github.com/PullRequestInc/go-gpt3"
)

type IGPT interface {
	ToGPT(text string) (*string, error)
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

func (gpt *igpt) ToGPT(text string) (*string, error) {

	c := *gpt.gptClint
	ctx := *gpt.ctx

	resp, err := c.ChatCompletionStream() CompletionStreamWithEngine(ctx, gpt3.TextDavinci003Engine, gpt3.CompletionRequest{
		Prompt: []string{
			"The first thing you should know about golang is",
		},
		MaxTokens: gpt3.IntPtr(30),
		Stop:      []string{"."},
		Echo:      true,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(resp.Choices[0].Text)

	a := (*resp).Object

	return &a, err
}
