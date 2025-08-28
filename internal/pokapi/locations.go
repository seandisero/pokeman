package pokapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreas(url *string) (ShallowLocationArea, error) {
	locationAreaPath := baseLocationAreaURL
	if url != nil {
		locationAreaPath = *url
	}

	if cachedLocation, exists := c.cache.Get(locationAreaPath); exists {

		locationAreas := ShallowLocationArea{}
		err := json.Unmarshal(cachedLocation, &locationAreas)
		if err != nil {
			return ShallowLocationArea{}, fmt.Errorf("error: %w", err)
		}
		return locationAreas, nil
	}

	req, err := http.NewRequest("GET", locationAreaPath, nil)
	if err != nil {
		return ShallowLocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ShallowLocationArea{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ShallowLocationArea{}, fmt.Errorf("error: %w", err)
	}

	locationAreas := ShallowLocationArea{}
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return ShallowLocationArea{}, fmt.Errorf("error: %w", err)
	}

	c.cache.Add(locationAreaPath, body)
	return locationAreas, nil
}

func (c *Client) GetLocationData(location string) (LocationData, error) {
	locationURL := baseLocationAreaURL + "/" + location

	if cached, exists := c.cache.Get(locationURL); exists {
		area := LocationData{}
		err := json.Unmarshal(cached, area)
		if err != nil {
			return LocationData{}, fmt.Errorf("error getting location data: %w", err)
		}
		return area, nil
	}

	req, err := http.NewRequest("GET", locationURL, nil)
	if err != nil {
		return LocationData{}, fmt.Errorf("error creating request GET %s: %w", locationURL, err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationData{}, fmt.Errorf("error getting location from %s: %w", locationURL, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationData{}, fmt.Errorf("error reading body %w", err)
	}

	var locationData LocationData
	err = json.Unmarshal(body, &locationData)
	if err != nil {
		return LocationData{}, fmt.Errorf("error unmarshaling data %w", err)
	}

	c.cache.Add(locationURL, body)
	return locationData, nil

}

func (c *Client) PrintCache() {
	for k, v := range c.cache.Entries {
		fmt.Println(k, v)
	}
}
