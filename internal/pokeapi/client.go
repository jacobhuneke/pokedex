package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jacobhuneke/pokedex/internal/pokecache"
)

type Client struct {
	Cache pokecache.Cache
}

func (c *Client) OpenFile(url string) ([]byte, error) {
	data, ok := c.Cache.Get(url)

	if ok {
		return data, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}

	c.Cache.Add(url, body)
	return body, nil
}

func GetDataAreaLoc(body []byte) (LocationAreaData, error) {
	data := LocationAreaData{}
	e := json.Unmarshal(body, &data)
	if e != nil {
		return LocationAreaData{}, e
	}
	return data, nil
}

func GetDataLoc(body []byte) (LocationData, error) {
	data := LocationData{}
	e := json.Unmarshal(body, &data)
	if e != nil {
		return LocationData{}, e
	}
	return data, nil
}
