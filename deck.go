package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades,", "Diamods", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two,", "Three", "Four"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuit)
		}
	}
	return cards
}

// reciver function
// any variable of type deck has access to this method
// reciver method has one letter variables
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), "")
}

func (d deck) shufleDeck() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) saveToFile(filename string) error {
	fmt.Printf("Saving deck %v to file %v\n", d, filename)
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	fmt.Printf("Reading deck from file: %v\n", filename)
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Log the error and entirly quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	slice := strings.Split(string(bs), ",")
	return deck(slice)
}
