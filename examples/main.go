package main

import (
	"fmt"
	"github.com/devalexandre/langflowgo/langflowclient"
	"log"
)

func main() {
	flowIdOrName := "b87811d9-9974-4d11-bfab-4d4ff0b43160"
	inputValue := "Qual o CNAE usado para pesca?"
	stream := false
	langflowClient := langflowclient.NewLangflowClient()
	tweaks := langflowclient.Options{
		"OpenAIEmbeddings-gMvoo": langflowclient.Options{},
		"Qdrant-NRPl4":           langflowclient.Options{},
		"SplitText-KUnLN":        langflowclient.Options{},
		"File-kgHNu":             langflowclient.Options{},
		"Qdrant-vpLtI":           langflowclient.Options{},
		"ChatInput-oYQ3r":        langflowclient.Options{},
		"ParseData-XlJWG":        langflowclient.Options{},
		"OpenAIEmbeddings-5zkfa": langflowclient.Options{},
		"Prompt-B8Yop":           langflowclient.Options{},
		"OpenAIModel-V3jzS":      langflowclient.Options{},
		"ChatOutput-1TK8r":       langflowclient.Options{},
	}

	response, err := langflowClient.RunFlow(
		flowIdOrName,
		inputValue,
		tweaks,
		stream,
		func(data map[string]interface{}) {
			fmt.Println("Received:", data)
		},
		func(message string) {
			fmt.Println("Stream Closed:", message)
		},
		func(err error) {
			fmt.Println("Stream Error:", err)
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	if !stream {
		if len(response.Outputs) > 0 {
			firstOutput := response.Outputs[0]
			if len(firstOutput.Outputs) > 0 {
				outputDetails := firstOutput.Outputs[0]
				if message := outputDetails.Results.Message.Data.Text; message != "" {
					fmt.Println("Final Output:", message)
				}
			}
		}
	}
}
