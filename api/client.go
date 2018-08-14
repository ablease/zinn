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

func (c *Client) Professions() ([]string, error) {
	fullURL := c.URL + "/v2/professions"
	req, err := http.NewRequest("GET", fullURL, nil)
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

	var profs []string
	err = json.Unmarshal(body, &profs)
	if err != nil {
		return nil, err
	}

	return profs, nil
}

func (c *Client) Masteries() ([]int, error) {
	fullURL := c.URL + "/v2/masteries"
	req, err := http.NewRequest("GET", fullURL, nil)
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

	var masts []int
	err = json.Unmarshal(body, &masts)
	if err != nil {
		return nil, err
	}

	return masts, nil
}
