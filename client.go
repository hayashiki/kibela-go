package kibela

import (
	"log"
	"net/http"
	"net/url"

	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const (
	defaultBaseURL = "https://%s.kibe.la/api/v1"
	envKibelaTOKEN = "KIBELA_TOKEN"
)

type Client struct {
	BaseURL     *url.URL
	AccessToken string
	Client      *http.Client

	Note *NoteService
}

type Payload struct {
	Query     string      `json:"query"`
	Variables interface{} `json:"variables,omitempty"`
}

type Errors []error

type Response struct {
	Errors Errors          `json:"errors,omitempty"`
	Data   json.RawMessage `json:"data,omitempty"`
}

type SuccessResponse struct {
	Success bool `json:"success"`
}

func NewClient(httpClient *http.Client, team string) (*Client, error) {

	accessToken := os.Getenv(envKibelaTOKEN)

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, err := url.Parse(fmt.Sprintf(defaultBaseURL, team))

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	c := &Client{
		BaseURL:     baseURL,
		AccessToken: accessToken,
		Client:      httpClient,
	}

	c.Note = &NoteService{client: c}

	return c, nil
}

func (c *Client) NewRequest(method string, body interface{}) (*http.Request, error) {
	u, err := url.Parse(c.BaseURL.String())
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	log.Printf("log body %v", body)

	if body != nil {
		b := new(bytes.Buffer)
		err := json.NewEncoder(b).Encode(body)
		if err != nil {
			return nil, err
		}
		buf = b
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "hayashiki")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	return req, nil
}

func (c *Client) Do(r *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.Client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&v)

	if err != nil {
		return resp, err
	}
	return resp, nil
}
