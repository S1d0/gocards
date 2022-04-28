package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type DadJokeRespone struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func main() {

	fmt.Println("Calling API ...")
	get()
}

func get() {
	fmt.Println("Calling API ...")
	resp, err := http.Get("https://icanhazdadjoke.com/")

	if err != nil {
		log.Fatalln(err)
	}

	// Defer ensures that the resp.Body.Close() is executed at the end of the method
	// Without it, possible resource leaks
	defer resp.Body.Close()

	responseBytes, _ := ioutil.ReadAll(resp.Body)
	var dadJokeResponse DadJokeRespone
	json.Unmarshal(responseBytes, &dadJokeResponse)

	fmt.Println(dadJokeResponse)
}
