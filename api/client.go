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

func getList(url string) ([]string, error) {
	body, err := get(url)
	if err != nil {
		return nil, err
	}

	var result []string
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func getIDs(url string) ([]int, error) {
	body, err := get(url)
	if err != nil {
		return nil, err
	}

	var result []int
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) Professions() ([]string, error) {
	return getList(c.URL + "/v2/professions")
}

func (c *Client) DailyCrafting() ([]string, error) {
	return getList(c.URL + "/v2/dailycrafting")
}

func (c *Client) MapChests() ([]string, error) {
	return getList(c.URL + "/v2/mapchests")
}
