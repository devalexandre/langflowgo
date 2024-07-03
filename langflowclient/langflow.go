package langflowclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Options map[string]interface{}

type LangflowClient struct {
	BaseURL string
	APIKey  string
}

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

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("HTTP error! Status: %d", resp.StatusCode))
	}

	return ioutil.ReadAll(resp.Body)
}

func (client *LangflowClient) initiateSession(flowId string, inputValue string, stream bool, tweaks map[string]interface{}) ([]byte, error) {
	endpoint := fmt.Sprintf("/api/v1/run/%s?stream=%t", flowId, stream)
	body := map[string]interface{}{
		"input_value": inputValue,
		"tweaks":      tweaks,
	}
	return client.post(endpoint, body)
}

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

func (client *LangflowClient) RunFlow(flowIdOrName string, inputValue string, tweaks Options, stream bool, onUpdate func(map[string]interface{}), onClose func(string), onError func(error)) (ResponseRoot, error) {
	var initResponse ResponseRoot
	responseBytes, err := client.initiateSession(flowIdOrName, inputValue, stream, tweaks)
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
