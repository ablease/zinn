package api

import (
	"encoding/json"
	"strconv"
)

type Achievement struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Requirement string   `json:"requirement"`
	LockedText  string   `json:"locked_text"`
	Type        string   `json:"type"`
	Flags       []string `json:"flags"`
	Tiers       []struct {
		Count  int `json:"count"`
		Points int `json:"points"`
	} `json:"tiers"`
	Rewards []struct {
		Type  string `json:"type"`
		ID    int    `json:"id,omitempty"`
		Count int    `json:"count"`
	} `json:"rewards"`
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

// Achievements returns achievments for a specific set of ids
func (c *Client) Achievements(ids []int) ([]Achievement, error) {
	//construct url from achievement ids
	fullURL := c.URL + "?ids="
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

	var achievements []Achievement
	err = json.Unmarshal(body, &achievements)
	if err != nil {
		// TODO WHen Unmarshaling fails, can we return the json object? The api gives nice errors :)
		return nil, err
	}

	return achievements, nil
}
