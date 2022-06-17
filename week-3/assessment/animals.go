package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) Eat() string {
	return a.food
}

func (a *Animal) Move() string {
	return a.locomotion
}

func (a *Animal) Speak() string {
	return a.noise
}

func main() {

	in := bufio.NewReader(os.Stdin)

	const (
		kind = iota
		action
	)

	const (
		cow = iota
		bird
		snake
	)

	animals := []Animal{
		{
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		},
		{
			food:       "worms",
			locomotion: "fly",
			noise:      "peep",
		},
		{
			food:       "mice",
			locomotion: "slither",
			noise:      "hsss",
		},
	}

	for {

		fmt.Println(">")
		a, err := in.ReadString('\n')
		userInput := strings.Fields(a)
		animal := userInput[kind]
		verb := userInput[action]
		if err != nil {
			log.Fatal(err)
		}

		switch {
		case animal == "cow" && verb == "speak":
			fmt.Println(animals[cow].Speak())
		case animal == "cow" && verb == "eat":
			fmt.Println(animals[cow].Eat())
		case animal == "cow" && verb == "move":
			fmt.Println(animals[cow].Move())

		case animal == "bird" && verb == "speak":
			fmt.Println(animals[bird].Speak())
		case animal == "bird" && verb == "eat":
			fmt.Println(animals[bird].Eat())
		case animal == "bird" && verb == "move":
			fmt.Println(animals[bird].Move())

		case animal == "snake" && verb == "speak":
			fmt.Println(animals[snake].Speak())
		case animal == "snake" && verb == "eat":
			fmt.Println(animals[snake].Eat())
		case animal == "snake" && verb == "move":
			fmt.Println(animals[snake].Move())
		default:
			fmt.Println("Invalid option")
		}
	}
}
