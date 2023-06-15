package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Client struct {
	URL string
}

type Mastery struct {
	Name string `json:"name"`
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

func (c *Client) GetMasteryIDs() ([]int, error) {
	fullURL := c.URL + "/v2/masteries"
	body, err := get(fullURL)
	if err != nil {
		return nil, err
	}

	var masteryIDs []int
	err = json.Unmarshal(body, &masteryIDs)
	if err != nil {
		return nil, err
	}
	return masteryIDs, nil
}

func (c *Client) Masteries() ([]string, error) {
	masteryIDs, err := c.GetMasteryIDs()
	if err != nil {
		return nil, err
	}

	masteryNames := []string{}
	for _, id := range masteryIDs {
		s := strconv.Itoa(id)

		fullURL := c.URL + "/v2/masteries/" + s
		body, err := get(fullURL)
		if err != nil {
			return nil, err
		}

		var mastery Mastery
		err = json.Unmarshal(body, &mastery)
		if err != nil {
			return nil, err
		}

		masteryNames = append(masteryNames, mastery.Name)
	}

	return masteryNames, nil
}

// AchievementIDs returns a list of all achievement ids
func (c *Client) AchievementIDs() ([]int, error) {
	achievementIDs := []int{}
	body, err := get(c.URL)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &achievementIDs)
	if err != nil {
		return nil, err
	}

	return achievementIDs, nil
}
