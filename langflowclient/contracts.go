package langflowclient

import (
	"net/http"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type APIStatusResponse struct {
	Enabled bool `json:"enabled"`
}

type APIKeyCheckResponse struct {
	HasApiKey bool `json:"has_api_key"`
	IsValid   bool `json:"is_valid"`
}

type VersionResponse struct {
	Version string `json:"version"`
}

type ResponseRoot struct {
	SessionID string    `json:"session_id"`
	Outputs   []Outputs `json:"outputs"`
}

type Outputs struct {
	Outputs []InnerOutput `json:"outputs,omitempty"` // camada extra para stream=true
	Results *Results      `json:"results,omitempty"` // stream=false
}

type InnerOutput struct {
	Results *Results `json:"results,omitempty"`
}

type Results struct {
	Message Message `json:"message"`
}

type Message struct {
	Timestamp  string `json:"timestamp"`
	Sender     string `json:"sender"`
	SenderName string `json:"sender_name"`
	Text       string `json:"text"`
}
