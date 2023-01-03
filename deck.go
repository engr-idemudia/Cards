package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// creat a new type of deck
// which is a slice
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

// func (d deck) saveToFile(filename string) error {
// 	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
// }

// use os.WriteFile package - modern way
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	// bs, err := ioutil.ReadFile(filename)
	bs, err := os.ReadFile(filename)
	if err != nil {
		//option #1 - log the error and return a call to newDeck()
		//option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1) //option #2
	}
	s := strings.Split(string(bs), ", ") //type conversion from []byte to string
	return deck(s)                       // to deck
}
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i] //swap
	}
}
