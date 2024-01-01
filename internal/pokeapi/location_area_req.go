package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationArea(url *string) (LocationAreaResponse, error) {
	endpoint := "/location-area"
	fullURL := baseUrl + endpoint
	if url != nil {
		fullURL = *url
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

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationAreasResp := LocationAreaResponse{}
	err = json.Unmarshal(dat, &locationAreasResp)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	return locationAreasResp, nil
}
