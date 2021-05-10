package client

import (
	"net/http"
)

type Client struct {
	hostname   string
	port       int
	httpClient *http.Client
}

func NewClient(hostname string, port int) *Client {
	return &Client{
		hostname:   hostname,
		port:       port,
		httpClient: &http.Client{},
	}
}
