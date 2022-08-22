package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var command string
	// Creates slice with first 3 elements equals with 0
	var sortedIntegers = make([]int, 3)
	// Used when the length lower than capacity. Important for the first 3 elements in slice
	//var indexOfSortedSlice = 0
	for {
		fmt.Println("Write number and press Enter to continue program. \nWrite \"X\" character and press enter to exit program")
		fmt.Scan(&command)

		// If the user introduced X in the console then we stop the program
		if command == "X" {
			break
		}

		// Conversion to number
		var enteredNumber, err = strconv.Atoi(command)
		// If the conversion is unsucsessfull show error and skip this element
		if err != nil {
			fmt.Println("ERROR: " + err.Error())
			continue
		}

		sortedIntegers = append(sortedIntegers, enteredNumber)

		// Sorting slice in ascendign order
		sort.Slice(sortedIntegers, func(i, j int) bool {
			return sortedIntegers[i] < sortedIntegers[j]
		})

		fmt.Printf("sortedIntegers: %v\n", sortedIntegers)
	}
}
