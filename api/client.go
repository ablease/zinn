package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	URL string
}

func NewClient(url string) *Client {
	return &Client{
		URL: url,
	}
}

// get makes a HTTP GET request to the provided url and returns the response body as bytes
func get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	httpClient := http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}

func (c *Client) Professions() ([]string, error) {
	fullURL := c.URL + "/v2/professions"
	body, err := get(fullURL)
	if err != nil {
		return nil, err
	}

	var profs []string
	err = json.Unmarshal(body, &profs)
	if err != nil {
		return nil, err
	}

	return profs, nil
}
