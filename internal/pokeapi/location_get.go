package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationData(locationName string) (LocationData, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationResp := LocationData{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return LocationData{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationData{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationData{}, err
	}

	locationData := LocationData{}
	err = json.Unmarshal(dat, &locationData)
	if err != nil {
		return LocationData{}, err
	}

	return locationData, nil
}
