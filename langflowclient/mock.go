package langflowclient

import (
	"net/http"
)

// MockHttpClient is a mock of HttpClient for testing.
type MockHttpClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}
