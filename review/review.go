package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	sound      string
}

func initData(nameMap map[string]Animal) {
	nameMap["cow"] = Animal{"grass", "walk", "moo"}
	nameMap["bird"] = Animal{"worms", "fly", "peep"}
	nameMap["snake"] = Animal{"mice", "slither", "hss"}
}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.sound)
}

func executeQuery(nameMap map[string]Animal) {
	fmt.Println("Input request or type exit to quit the program: ")
	for {
		fmt.Printf(">")
		consoleReader := bufio.NewReader(os.Stdin)
		inputStr, _ := consoleReader.ReadString('\n')
		inputStr = strings.Trim(inputStr, "\n")
		inputList := strings.Split(inputStr, " ")
		lenInpList := len(inputList)
		switch lenInpList {
		case 0:
			fmt.Println("Please provide input request")
		case 1:
			if inputList[0] == "exit" {
				return
			} else {
				fmt.Println("Please provide correct input")
			}
		case 2:
			if (inputList[0] == "cow" || inputList[0] == "bird" || inputList[0] == "snake") && (inputList[1] == "eat" || inputList[1] == "move" || inputList[1] == "speak") {
				if inputList[1] == "eat" {
					nameMap[inputList[0]].Eat()
				} else if inputList[1] == "move" {
					nameMap[inputList[0]].Move()
				} else if inputList[1] == "speak" {
					nameMap[inputList[0]].Speak()
				}
			} else {
				fmt.Println("Please provide correct input")
			}
		default:
			fmt.Println("Please provide correct input")
		}
	}
}

func main() {
	nameMap := make(map[string]Animal)
	initData(nameMap)
	// fmt.Println(nameMap)
	executeQuery(nameMap)
}
