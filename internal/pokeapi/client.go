package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func OpenFile(url string) ([]byte, error) {
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
	return body, nil
}

func GetDataLoc(body []byte) (LocationData, error) {
	data := LocationData{}
	e := json.Unmarshal(body, &data)
	if e != nil {
		return LocationData{}, e
	}
	return data, nil
}
