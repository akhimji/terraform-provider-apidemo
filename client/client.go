package client

type APIClient struct {
	hostname string
	port     string
	key      string
}

func NewClient(hostname string, port string, key string) *APIClient {
	return &APIClient{
		hostname: hostname,
		port:     port,
		key:      key,
	}
}
