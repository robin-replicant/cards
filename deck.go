package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// we are going to create a new type, deck type
// which is a slice of strings

type deck []string

// we dont add the recievier coz we are creating a new deck
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"♠", "♥", "♦", "♣"}
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" "+suit)
		}
	}

	return cards
}

// we add a reciver on a fuction
// any variable of type deck now
// gets access to the methods that we set a receiver
// cards is the actual copy of the deck we are working with
// for convetion we dont use this or self to refeer to cards

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	//anyone can write and read this file
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		// option #1 - log the error and return a call to new Deck()
		// option #2 - log the error and quit the program
		// pick option #2
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := deck(strings.Split(string(byteSlice), ","))
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		// swipe
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
