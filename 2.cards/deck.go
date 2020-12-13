package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create type of 'deck' which is a slice of strings // deck is something like a class in OOP
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) deckToString() string {
	slice := []string(d)
	strings := strings.Join(slice, ", ")
	return strings
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.deckToString()), 0666)
}

func newDeckFromFile(filename string) deck {
	fmt.Println("Info: loading deck from file:", filename)
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option 1) - log the error and return a call to newDeck()
		// option 2) - Log the error and quit program
		fmt.Println("Error:", err)
		fmt.Println("Error: Quiting the program")
		os.Exit(1)
	}

	strCSV := string(bs)
	s := strings.Split(strCSV, ", ")
	return deck(s)
}

func (d deck) shuffleDeck() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		// newI := rand.Intn(len(d) - 1)
		newI := r.Intn(len(d) - 1)

		d[i], d[newI] = d[newI], d[i]
	}
}
