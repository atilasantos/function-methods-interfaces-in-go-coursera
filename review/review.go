package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Function BubbleSort(slice) - Sort a slice of integers in ascending order
func BubbleSort(slice []int) {
	// Check that the slice contains at least two elements
	if len(slice) > 1 {
		for imax := len(slice) - 1; imax > 0; imax -= 1 {
			for i := 0; i < imax; i += 1 {
				if slice[i] > slice[i+1] {
					Swap(slice, i)
				}
			}
		}
	}
}

// Swap i'th and (i+1)'th element of the slice using a temporary variable
func Swap(slice []int, idx int) {
	itmp := slice[idx]
	slice[idx] = slice[idx+1]
	slice[idx+1] = itmp
}

func main() {
	// Create an empty slice
	sli := make([]int, 0)

	// Ask user for the list of numbers
	fmt.Print("Type a sequence of integers separated by spaces: ")
	sread, err := bufio.NewReader(os.Stdin).ReadString('\n') // Read a string from stdin
	if err == nil {
		slist := strings.Split(strings.Trim(sread, " \r\n"), " ") // Split trimmed string into slice of numbers (as strings)
		for _, sval := range slist {                              // For every number (as string) in the slice ...
			ival, err := strconv.Atoi(sval) // Convert to integer
			if err == nil {
				sli = append(sli, ival) // Append to slice of integers if conversion was successful
			}
		}
	} else {
		return
	}

	// Sort slice
	BubbleSort(sli)

	// Print sorted slice
	fmt.Print("Sorted list of values: ")
	for _, ival := range sli {
		fmt.Printf("%d ", ival)
	}
	fmt.Println()
}
