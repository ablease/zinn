// Package gw2api parses JSON responses from the gw2 api into usable, exported
// go structs. In order to do this, gw2api provides an http client and
// a mechanism to provide authentication information for the http client
// to use against the authenticated endpoints of the gw2 api
package gw2api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Gw2Api struct {
	url    string
	client http.Client
	auth   string
}

// NewAPI returns a new Gw2Api with an instantiated http client
// the url to be used by the client must be injectable!!
func NewAPI(url string) *Gw2Api {
	if url == "" {
		url = "https://api.guildwars2.com"
	}

	api := &Gw2Api{
		client: http.Client{},
		url:    url,
	}
	return api
}

// NewAuthenticatedAPI()

type Mastery struct {
	Name string `json:"name"`
}

func (gw2api *Gw2Api) Professions() ([]string, error) {
	fullURL := gw2api.url + "/v2/professions"
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

func (gw2api *Gw2Api) GetMasteryIDs() ([]int, error) {
	fullURL := gw2api.url + "/v2/masteries"
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

	var masteryIDs []int
	err = json.Unmarshal(body, &masteryIDs)
	if err != nil {
		return nil, err
	}
	return masteryIDs, nil
}

func (gw2api *Gw2Api) Masteries() ([]Mastery, error) {
	masteryIDs, err := gw2api.GetMasteryIDs()
	if err != nil {
		return nil, err
	}

	masterys := []Mastery{}
	for _, id := range masteryIDs {
		s := strconv.Itoa(id)

		fullURL := gw2api.url + "/v2/masteries/" + s
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

		var mastery Mastery
		err = json.Unmarshal(body, &mastery)
		if err != nil {
			return nil, err
		}

		masterys = append(masterys, mastery)
	}

	return masterys, nil
}
