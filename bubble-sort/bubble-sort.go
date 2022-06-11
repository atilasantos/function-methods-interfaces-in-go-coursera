package main

import (
	"fmt"
	"strconv"
)

func Swap(intItems []int, swapItem int) {
	intItems[swapItem], intItems[swapItem+1] = intItems[swapItem+1], intItems[swapItem]

}

func BubbleSort(items []int) {
	itemsSize := len(items)
	for i := 0; i < itemsSize-1; i++ {
		for j := 0; j < itemsSize-1; j++ {
			if items[j] > items[j+1] {
				Swap(items, j)
			}
		}
	}
}

func main() {
	var numbers []int
	var number string
	var n int

	for n < 10 {
		fmt.Println("Type an integer number: ")
		_, err := fmt.Scan(&number)
		nInt, _ := strconv.Atoi(number)
		if err != nil {
			fmt.Print(err)
		}
		numbers = append(numbers, nInt)
		n++
	}
	fmt.Println(numbers)
	BubbleSort(numbers)
	fmt.Println(numbers)

}
