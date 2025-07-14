package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	respData, exists := c.cache.Get(url)
	if !exists {

		resp, err := http.Get(url)
		if err != nil {
			return RespShallowLocations{}, err
		}
		defer resp.Body.Close()
		dat, err := io.ReadAll(resp.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}
		respData = dat
		c.cache.Add(url, respData)
	}

	var locationsResp RespShallowLocations
	err := json.Unmarshal(respData, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil

}
