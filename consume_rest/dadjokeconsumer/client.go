package dadjokeconsumer

import (
	"net/http"
	"net/url"
)

type Client struct {
	BaseUrl    *url.URL
	httpClient *http.Client
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