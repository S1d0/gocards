package dadjokeconsumer

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
)

func (c *Client) GetJoke() DadJoke {
	req := c.newRequest("GET", "/", nil)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	// Make sure that body will get closed
	defer resp.Body.Close()
	var dadJoke DadJoke
	err = json.NewDecoder(resp.Body).Decode(&dadJoke)

	if err != nil {
		log.Fatalln(err)
	}

	return dadJoke
}

// helper method
func (c *Client) newRequest(method string, path string, body io.Reader) *http.Request {
	endpoint := &url.URL{Path: path}
	fullPath := c.BaseUrl.ResolveReference(endpoint)

	req, err := http.NewRequest(method, fullPath.String(), body)
	if err != nil{ 
		log.Fatalln("Creating request failed, error:",err)
	}
	req.Header.Add("Accept", "application/json")

	return req
}
