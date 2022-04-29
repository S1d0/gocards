package dadjokeconsumer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Public class
type DadJoke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func ConsumeJoke() DadJoke {
	url := "https://icanhazdadjoke.com/"
	client := http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/json")
	
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var dadJokeResponse DadJoke
	json.Unmarshal(bodyBytes, &dadJokeResponse)

	return dadJokeResponse
}
