package langflowclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Options map for additional configurations
type Options map[string]interface{}

// LangflowClient structure
type LangflowClient struct {
	BaseURL    string
	APIKey     string
	HttpClient HttpClient
}

// NewLangflowClient creates a new LangflowClient with optional configurations
func NewLangflowClient(options ...Option) *LangflowClient {
	client := &LangflowClient{
		BaseURL:    "http://127.0.0.1:7860",
		APIKey:     "",
		HttpClient: http.DefaultClient,
	}
	for _, option := range options {
		option(client)
	}
	return client
}

// Request and Response structs
type RunFlowRequest struct {
	InputValue string                 `json:"input_value"`
	Tweaks     map[string]interface{} `json:"tweaks"`
}

type RunFlowResponse struct {
	Outputs []struct {
		Artifacts struct {
			Type      string `json:"type"`
			StreamURL string `json:"stream_url"`
		} `json:"artifacts"`
	} `json:"outputs"`
}

// post sends a POST request
func (client *LangflowClient) post(endpoint string, body interface{}) ([]byte, error) {
	url := client.BaseURL + endpoint
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if client.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+client.APIKey)
	}

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("HTTP error! Status: %d", resp.StatusCode))
	}

	return ioutil.ReadAll(resp.Body)
}

// get sends a GET request
func (client *LangflowClient) get(endpoint string) ([]byte, error) {
	url := client.BaseURL + endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if client.APIKey != "" {
		req.Header.Set("Authorization", "Bearer "+client.APIKey)
	}

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("HTTP error! Status: %d", resp.StatusCode))
	}

	return ioutil.ReadAll(resp.Body)
}

// initiateSession initiates a new session
func (client *LangflowClient) initiateSession(flowId string, request RunFlowRequest) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/v1/run/%s?stream=%t", flowId, false)
	return client.post(endpoint, request)
}

// handleStream handles a streaming response
func (client *LangflowClient) handleStream(streamUrl string, onUpdate func(map[string]interface{}), onClose func(string), onError func(error)) {
	resp, err := http.Get(streamUrl)
	if err != nil {
		onError(err)
		return
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	for {
		var data map[string]interface{}
		if err := decoder.Decode(&data); err != nil {
			onError(err)
			break
		}
		onUpdate(data)
	}

	onClose("Stream closed")
}

// RunFlow runs a flow
func (client *LangflowClient) RunFlow(flowIdOrName string, inputValue string, tweaks Options, stream bool, onUpdate func(map[string]interface{}), onClose func(string), onError func(error)) (ResponseRoot, error) {
	var initResponse ResponseRoot
	request := RunFlowRequest{
		InputValue: inputValue,
		Tweaks:     tweaks,
	}
	responseBytes, err := client.initiateSession(flowIdOrName, request)
	if err != nil {
		onError(err)
		return initResponse, err
	}
	err = json.Unmarshal(responseBytes, &initResponse)
	if err != nil {
		onError(err)
		return initResponse, err
	}

	// Check if streaming is requested and process accordingly
	if stream && len(initResponse.Outputs) > 0 {
		for _, output := range initResponse.Outputs {
			if output.Artifacts.Type == "stream" && output.Artifacts.StreamURL != "" {
				fullStreamURL := client.BaseURL + output.Artifacts.StreamURL
				fmt.Println("Streaming from:", fullStreamURL)
				go client.handleStream(fullStreamURL, onUpdate, onClose, onError)
				break
			}
		}
	}

	return initResponse, nil
}

// New methods
func (client *LangflowClient) CheckAPIStatus() (APIStatusResponse, error) {
	var result APIStatusResponse
	responseBytes, err := client.get("/api/v1/store/check/")
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(responseBytes, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (client *LangflowClient) CheckAPIKey() (APIKeyCheckResponse, error) {
	var result APIKeyCheckResponse
	responseBytes, err := client.get("/api/v1/store/check/api_key")
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(responseBytes, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (client *LangflowClient) GetVersion() (VersionResponse, error) {
	var result VersionResponse
	responseBytes, err := client.get("/api/v1/version")
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(responseBytes, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (client *LangflowClient) GetAllFlows() (AllFlowsResponse, error) {
	var result AllFlowsResponse
	responseBytes, err := client.get("/api/v1/all")
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(responseBytes, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (client *LangflowClient) GetBuildStatus(buildID string) (BuildStatusResponse, error) {
	var result BuildStatusResponse
	endpoint := fmt.Sprintf("/api/v1/build/%s/status", buildID)
	responseBytes, err := client.get(endpoint)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(responseBytes, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
