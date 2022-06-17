package main

import "fmt"

func applyIt(afunct func(int) int, val int) int {
	return afunct(val)
}

func sumWithTen(number int) int {
	return number + 10
}

func subtractWithFive(number int) int {
	return number - 5
}

func totalSum(intSlice ...int) int {
	var total int
	for _, v := range intSlice {
		total = total + v
	}
	return total
}

func main() {
	// passing function as argument
	fmt.Println(applyIt(sumWithTen, 10))
	fmt.Println(applyIt(subtractWithFive, 10))
	//anonymous function
	fmt.Println(applyIt(func(x int) int { return x * 3 }, 10))
	//variadic function
	fmt.Println(totalSum(10, 20, 30, 40))

}
