package pokeapi

import (
	"encoding/json"
	"fmt"
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

func (c *Client) GetLocation(name string) (Location, error) {
	url := baseURL + "/location-area/" + name

	respData, exists := c.cache.Get(url)
	if !exists {

		resp, err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			return Location{}, fmt.Errorf("Request failed to : %s and returned: %s", url, err)
		}
		if resp.StatusCode > 299 {
			return Location{}, fmt.Errorf("Request failed with error code: %d", resp.StatusCode)
		}
		dat, err := io.ReadAll(resp.Body)
		if err != nil {
			return Location{}, fmt.Errorf("Failed to read response body: %v and returned: %w", resp.Body, err)
		}
		respData = dat
		c.cache.Add(url, dat)

	}
	var locationResp Location
	err := json.Unmarshal(respData, &locationResp)
	if err != nil {
		return Location{}, fmt.Errorf("Failed to parse json from: %s returned: %s", url, err)
	}
	return locationResp, nil
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	respData, exists := c.cache.Get(url)
	if !exists {
		resp, err := http.Get(url)
		if err != nil {
			return Pokemon{}, fmt.Errorf("Request to: %s failed with: %w", url, err)
		}
		defer resp.Body.Close()
		if resp.StatusCode > 299 {
			return Pokemon{}, fmt.Errorf("Request failed with the error code: %d", resp.StatusCode)
		}
		dat, err := io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, fmt.Errorf("Failed to read response body: %v and returned: %w", resp.Body, err)
		}
		respData = dat
		c.cache.Add(url, dat)
	}
	var pokemon Pokemon
	err := json.Unmarshal(respData, &pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Failed to parse json and returned: %s", err)
	}
	return pokemon, nil
}
