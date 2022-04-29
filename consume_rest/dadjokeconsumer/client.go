package dadjokeconsumer

import (
	"net/http"
	"net/url"
)

type Client struct {
	BaseUrl    *url.URL
	httpClient *http.Client
}

type DadJoke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func NewClient() *Client {
	baseUrl := url.URL{
		Scheme: "https",
		Host:   "icanhazdadjoke.com",
		Path:   "/",
	}

	httpClient := http.Client{}
	return &Client{
		BaseUrl:    &baseUrl,
		httpClient: &httpClient,
	}

}
