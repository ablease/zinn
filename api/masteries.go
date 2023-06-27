package api

import (
	"encoding/json"
	"strconv"
)

type Mastery struct {
	ID          int     `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Requirement string  `json:"requirement,omitempty"`
	Order       int     `json:"order,omitempty"`
	Background  string  `json:"background,omitempty"`
	Region      string  `json:"region,omitempty"`
	Levels      []Level `json:"levels,omitempty"`
}

type Level struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Instruction string `json:"instruction,omitempty"`
	Icon        string `json:"icon,omitempty"`
	PointCost   int    `json:"point_cost,omitempty"`
	ExpCost     int    `json:"exp_cost,omitempty"`
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

func (c *Client) Masteries(ids []int) ([]Mastery, error) {
	//construct url from mastery ids
	fullURL := c.URL + "/v2/masteries?ids="
	for _, id := range ids {
		if ids[len(ids)-1] == id {
			fullURL = fullURL + strconv.Itoa(id)
			break
		}
		fullURL = fullURL + strconv.Itoa(id) + ","
	}

	body, err := get(fullURL)
	if err != nil {
		return nil, err
	}

	var masteries []Mastery
	err = json.Unmarshal(body, &masteries)
	if err != nil {
		// TODO WHen Unmarshaling fails, can we return the json object? The api gives nice errors :)
		return nil, err
	}

	return masteries, nil
}
