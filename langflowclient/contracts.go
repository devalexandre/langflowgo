package langflowclient

import (
	"time"
)

// ResponseRoot represents the top-level structure of the response from Langflow API.
type ResponseRoot struct {
	SessionID string    `json:"session_id"`
	Outputs   []Output  `json:"outputs"`
	Messages  []Message `json:"messages"`
}

// Output represents the "outputs" part of the response.
type Output struct {
	Inputs    Input           `json:"inputs"`
	Outputs   []OutputDetails `json:"outputs"`
	Artifacts Artifacts       `json:"artifacts"`
}

// Input represents details about the input provided to the flow.
type Input struct {
	InputValue string `json:"input_value"`
}

// OutputDetails represents the details within each output, often containing results and messages.
type OutputDetails struct {
	Results Result `json:"results"`
	Type    string `json:"type"`
}

// Result represents the detailed result information.
type Result struct {
	Message Message `json:"message"`
}

// Message represents a message within results or top-level messages.
type Message struct {
	TextKey    string    `json:"text_key"`
	Data       TextData  `json:"data"`
	Sender     string    `json:"sender"`
	SenderName string    `json:"sender_name"`
	SessionID  string    `json:"session_id"`
	Timestamp  time.Time `json:"timestamp"` // Adjusted to parse correctly
	Files      []File    `json:"files"`
	FlowID     string    `json:"flow_id"`
}

// TextData represents the actual text data sent by the machine.
type TextData struct {
	Text string `json:"text"`
}

// Artifacts represents additional data associated with a response.
type Artifacts struct {
	Message    string `json:"message"`
	Sender     string `json:"sender"`
	SenderName string `json:"sender_name"`
	StreamURL  string `json:"stream_url"` // Ensure this line is correctly added
	Files      []File `json:"files"`
	Type       string `json:"type"`
}

// File represents a file that may be attached to a message.
type File struct {
	// File fields here
}
