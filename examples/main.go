package main

import (
	"fmt"
	"github.com/devalexandre/langflowgo/langflowclient"
	"log"
)

func main() {
	flowIdOrName := ""
	inputValue := "Boa tarde"
	stream := true
	langflowClient := langflowclient.NewLangflowClient(
		langflowclient.WithHost(""),
		langflowclient.WithAPIKey(""),
	)
	tweaks := langflowclient.Options{}

	response, err := langflowClient.RunFlow(
		flowIdOrName,
		inputValue,
		tweaks,
		stream,
		func(err error) {
			fmt.Println("Stream Error:", err)
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	msg := langflowclient.GetLastMessage(response)

	fmt.Println("Message:", msg.Text)

}
