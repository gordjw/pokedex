package pokeapi

import (
	"io"
	"log"
	"net/http"
	"pokecache"
	"time"
)

const baseUrl = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(),
	}
}

func (c *Client) getPokemonApiData(url string) ([]byte, error) {
	bytes, ok := c.cache.Get(url)
	if ok {
		return bytes, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	bytes, err = io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	c.cache.Add(url, bytes)

	return bytes, nil
}
