package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Speak()
	Move()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (c Cow) Eat() {
	fmt.Println(c.food)
}

func (c Cow) Speak() {
	fmt.Println(c.noise)
}

func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (b Bird) Eat() {
	fmt.Println(b.food)
}

func (b Bird) Speak() {
	fmt.Println(b.noise)
}

func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (s Snake) Eat() {
	fmt.Println(s.food)
}

func (s Snake) Speak() {
	fmt.Println(s.noise)
}

func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

func animalExists(animals map[string]Animal, name string) bool {
	if _, ok := animals[name]; ok {
		return true
	}
	return false
}

func main() {

	input := bufio.NewReader(os.Stdin)

	const (
		cow = iota
		bird
		snake
	)

	const (
		action = iota
		name
		transactionType
	)

	animals := []Animal{
		Cow{
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		},
		Bird{
			food:       "worms",
			locomotion: "fly",
			noise:      "peep",
		},
		Snake{
			food:       "mice",
			locomotion: "slither",
			noise:      "hsss",
		},
	}

	newAnimals := make(map[string]Animal)
	newAnimals["mimosa"] = animals[cow]

	for {

		fmt.Print(">")
		line, err := input.ReadString('\n')
		splitedLine := strings.Fields(line)
		if err != nil {
			log.Fatal(err)
		}

		if len(splitedLine) == 3 {

			switch {
			case splitedLine[transactionType] == "cow" && splitedLine[action] == "newanimal":
				newAnimals[splitedLine[name]] = animals[cow]
			case splitedLine[transactionType] == "bird" && splitedLine[action] == "newanimal":
				newAnimals[splitedLine[name]] = animals[bird]
			case splitedLine[transactionType] == "snake" && splitedLine[action] == "newanimal":
				newAnimals[splitedLine[name]] = animals[snake]

			case splitedLine[transactionType] == "eat" && splitedLine[action] == "query":
				if animalExists(newAnimals, splitedLine[name]) {
					newAnimals[splitedLine[name]].Eat()
				} else {
					fmt.Println("Animal does not exist.")
				}

			case splitedLine[transactionType] == "speak" && splitedLine[action] == "query":
				if animalExists(newAnimals, splitedLine[name]) {
					newAnimals[splitedLine[name]].Speak()
				} else {
					fmt.Println("Animal does not exist.")
				}
			case splitedLine[transactionType] == "move" && splitedLine[action] == "query":
				if animalExists(newAnimals, splitedLine[name]) {
					newAnimals[splitedLine[name]].Move()
				} else {
					fmt.Println("Animal does not exist.")
				}
			default:
				fmt.Println("Invalid option")
			}
		} else {
			fmt.Println("Invalid number of parameters")
		}

	}
}
