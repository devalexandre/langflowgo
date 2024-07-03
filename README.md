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
    "log"

    "github.com/devalexandre/langflowgo/langflowclient"
)

func main() {
    client := langflowclient.LangflowClient{
        BaseURL: "http://127.0.0.1:7860",
        APIKey:  "your_api_key_here",
    }

    tweaks := langflowclient.Options{
        "OpenAIEmbeddings-gMvoo": {},
        // Add other tweaks as necessary
    }

    response, err := client.RunFlow("flow_id_here", "User message", tweaks, false,
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

    fmt.Println("Flow completed successfully:", response)
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

