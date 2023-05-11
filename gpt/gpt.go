package gpt

import (
	//"fmt"
	//"context"
	openai "github.com/sashabaranov/go-openai"
)

var GptClient = &openai.Client{}

func Connect(gptToken string) {
	GptClient = openai.NewClient(gptToken)

}
