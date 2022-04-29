package main

import (
	"fmt"
	"gocards/consume_rest/dadjokeconsumer"
)

func main() {
	dadJokeClient := dadjokeconsumer.NewClient()

	dadJoke := dadJokeClient.GetJoke()

	fmt.Println("Call status", dadJoke.Status)
	fmt.Println("Call ID", dadJoke.ID)
	fmt.Println("Joke:", dadJoke.Joke)
}
