# LangflowClient

The `LangflowClient` is a Go library designed to interact seamlessly with the Langflow API. It provides robust tools to run flows, handle streaming data, and retrieve outputs efficiently, making it ideal for applications requiring real-time data processing and response generation.

## Features

- **Run Flows**: Initiate and manage flows with custom input values and configurations.
- **Real-Time Streaming**: Handle real-time data streaming for immediate response processing.
- **Error Handling**: Comprehensive error handling to manage and troubleshoot API interaction effectively.
- **Flexible Tweaking**: Dynamic tweaking options to customize the flow execution.

## Getting Started

### Prerequisites

- Go 1.15 or higher
- Access to Langflow API with a valid API key

### Installation

```bash
go get github.com/devalexandre/langflowgo
```

# How To

Here is a simple example of how to use the client to run a flow:

```go
package main

import (
	"fmt"
	"github.com/devalexandre/langflowgo/langflowclient"
	"log"
)

func main() {
	flowIdOrName := "cbed7afd-328d-4eb1-a79f-348852ab7159"
	inputValue := "Boa tarde"
	stream := true
	langflowClient := langflowclient.NewLangflowClient(
		langflowclient.WithHost("http://192.168.63.17:7860"),
		langflowclient.WithAPIKey("sk-9l-7zxtsqDaFh0zQ3Ztvco7kz2DBHy8S4j-trF71m8Q"),
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
```

## Configuration
Configure the client with your specific API key and base URL. Ensure you handle the credentials securely.

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are greatly appreciated.

1. Fork the Project
2. Create your Feature Branch (git checkout -b feature/AmazingFeature)
3. Commit your Changes (git commit -m 'Add some AmazingFeature')
4. Push to the Branch (git push origin feature/AmazingFeature)
5. Open a Pull Request

## License
Distributed under the MIT License. See `LICENSE` for more information.

## Contact
Alexandre E Souza - alexandre@dev2learn.com
linkedin - https://www.linkedin.com/in/devevantelista/


### Explanation

- **Features**: Highlights the capabilities of your client.
- **Getting Started**: Provides step-by-step instructions on how to install and use the client.
- **Usage**: A basic example that demonstrates the client's functionality.
- **Contributing**: Encourages others to contribute to the project.
- **License**: Specifies the type of license under which the project is released.
- **Contact**: Provides contact information and a link to the project repository.

