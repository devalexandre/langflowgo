package langflowclient

// Option functional option type
type Option func(*LangflowClient)

// WithHost sets the base URL for the client
func WithHost(host string) Option {
	return func(c *LangflowClient) {
		c.BaseURL = host
	}
}

// WithAPIKey sets the API key for the client
func WithAPIKey(apiKey string) Option {
	return func(c *LangflowClient) {
		c.APIKey = apiKey
	}
}

// WithHttpClient sets the HTTP client for the client
func WithHttpClient(httpClient HttpClient) Option {
	return func(c *LangflowClient) {
		c.HttpClient = httpClient
	}
}
