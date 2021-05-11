package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"io"
	"encoding/json"

	"github.com/cavisd7/terraform-provider-scooter/api/server"
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

/*func (c *Client) GetItem(name string) (*server.Item, error) {

}*/

func (c *Client) NewItem(item *server.Item) error {
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(item)
	if err != nil {
		return err
	}

	_, err = c.httpRequest("item", "POST", buf)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) httpRequest(path, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	requestPath := fmt.Sprintf("%s:%v/%s", c.hostname, c.port, path)
	req, err := http.NewRequest(method, requestPath, &body)
	if err != nil {
		return nil, err 
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("non 200 status code: %v", resp.StatusCode)
		}

		return nil, fmt.Errorf("non 200 status code: %v - %s", resp.StatusCode, respBody.String())
	}

	return resp.Body, nil
}