package client

type APIClient struct {
	hostname string
	port     string
	key      string
}

func CreateClient(hostname string, port string, key string) *APIClient {
	return &APIClient{
		hostname: hostname,
		port:     port,
		key:      key,
	}
}

func (c *APIClient) GetConnString() string {
	baseurl := c.hostname + ":" + c.port + c.key
	return baseurl
}
