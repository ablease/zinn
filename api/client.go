package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ZinnClient struct {
	URL string
}

func NewZinnClient(url string) *ZinnClient {
	return &ZinnClient{
		URL: url,
	}
}

func (c *ZinnClient) Professions() ([]string, error) {
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
