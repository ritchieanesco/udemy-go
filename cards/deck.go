package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

// when you see a function with the params
// in front of the function name
// then this is called a receiver
// which means the function is now available
// for the that type
// in this instance you call the following function
// card.print() // card is declared from the type

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardSuite := range cardSuites {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuite)
		}
	}
	return cards
}

// split deck into two
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// receives a string slice and joins all the strings into a single string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// 0666 permissions (anyone can read or write)
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// bs = byte slice, err = error
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// log error
		fmt.Println("Error:", err)
		// exit program
		os.Exit(1)
	}
	// convert byte slice to string, then split string by ,
	s := strings.Split(string(bs), ",")

	// convert slice string to deck
	return deck(s)
}

func (d deck) shuffle() {

	// create new seed to ensure randomness
	// rand.NewSource expects int64
	// time unix nano provides int64 value

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// i = index
	for i := range d {
		// len is length of slice
		newPosition := r.Intn(len(d) - 1)
		// swap position of card
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
