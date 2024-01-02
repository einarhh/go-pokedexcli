package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationArea(url *string) (LocationAreaResponse, error) {
	endpoint := "/location-area?offset=0&limit=20"
	fullURL := baseUrl + endpoint
	if url != nil {
		fullURL = *url
	}

	// Check cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		locationAreasResp := LocationAreaResponse{}
		err := json.Unmarshal(dat, &locationAreasResp)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		return locationAreasResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaResponse{}, fmt.Errorf("Response failed with status code: %d", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationAreasResp := LocationAreaResponse{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreasResp, nil
}

func (c *Client) ExploreLocationArea(name string) (LocationAreaData, error) {
	endpoint := "/location-area/" + name
	fullURL := baseUrl + endpoint

	// Check cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		locationAreaData := LocationAreaData{}
		err := json.Unmarshal(dat, &locationAreaData)
		if err != nil {
			return LocationAreaData{}, err
		}

		return locationAreaData, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreaData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreaData{}, fmt.Errorf("Response failed with status code: %d", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaData{}, err
	}

	locationAreaData := LocationAreaData{}
	err = json.Unmarshal(dat, &locationAreaData)
	if err != nil {
		return LocationAreaData{}, err
	}

	c.cache.Add(fullURL, dat)

	return locationAreaData, nil
}
