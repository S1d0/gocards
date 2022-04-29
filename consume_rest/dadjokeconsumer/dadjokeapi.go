package dadjokeconsumer

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func (c *Client) GetJoke() DadJoke {
	rel := &url.URL{Path: "/"}
	u := c.BaseUrl.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatalln(err)
	}

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
