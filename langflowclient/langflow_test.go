package langflowclient_test

import (
	"fmt"
	"testing"

	"github.com/devalexandre/langflowgo/langflowclient"
)

func TestLangflowClient_RunFlow(t *testing.T) {
	flowIdOrName := "786fd3ec-b741-42ca-9c79-906edfe65851"
	inputValue := "what's Yamask attack?"
	stream := false
	langflowClient := langflowclient.NewLangflowClient()
	tweaks := langflowclient.Options{}

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
		t.Fatalf("Expected no error, got %v", err)
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
func TestLangflowClient_CheckAPIStatus(t *testing.T) {
	client := langflowclient.NewLangflowClient()

	res, err := client.CheckAPIStatus()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !res.Enabled {
		t.Errorf("Expected enabled to be true, got %v", res.Enabled)
	}
}

func TestLangflowClient_CheckAPIKey(t *testing.T) {
	t.Skip("Skipping test that requires an API key")
	client := langflowclient.NewLangflowClient(
		langflowclient.WithAPIKey("sk-EBNSQPg0xB5v95dTijlYJhGuTVbStfEo1izrnAScSzs"),
	)

	res, err := client.CheckAPIKey()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expectedDetail := "An API key as query or header, or a JWT token must be passed"
	if res.Detail != expectedDetail {
		t.Errorf("Expected detail to be '%s', got '%s'", expectedDetail, res.Detail)
	}
}

func TestLangflowClient_GetVersion(t *testing.T) {
	client := langflowclient.NewLangflowClient()

	res, err := client.GetVersion()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if res.Version == "" {
		t.Errorf("Expected a version, got empty string")
	}
}

func TestLangflowClient_GetAllFlows(t *testing.T) {
	client := langflowclient.NewLangflowClient()

	res, err := client.GetAllFlows()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if res.Chains.LLMChain.Template.Type == "" {
		t.Errorf("Expected at least one flow, got none")
	}
}

func TestLangflowClient_GetBuildStatus(t *testing.T) {
	t.Skip("Skipping test that requires a build ID")
	client := langflowclient.NewLangflowClient()

	buildID := "786fd3ec-b741-42ca-9c79-906edfe65851"
	res, err := client.GetBuildStatus(buildID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if res.Status == "" {
		t.Errorf("Expected a status, got empty string")
	}
}
