package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var scaner = bufio.NewScanner(os.Stdin)
	fmt.Print("Enter number of rows:")
	scaner.Scan()
	var numberOfRowsText = scaner.Text()
	var numberOfRowsInt, _ = strconv.Atoi(numberOfRowsText)

	fmt.Print("Enter element to print:")
	scaner.Scan()
	var symbolToPrint = scaner.Text()

	// var triangle = triangleLeft(numberOfRowsInt, symbolToPrint)
	var triangle = TriangleCenter(numberOfRowsInt, symbolToPrint)
	fmt.Print(triangle)
}

// The goal of this function is to print a triangle without spaces from the left.
// As input to the function are numberOfRows and symbolToPrint.
func TriangleLeft(numberOfRows int, symbolToPrint string) string {
	var finalString = ""
	for i := 1; i <= numberOfRows; i++ {
		var numberOfSymbols = 1 * i

		for j := 0; j < numberOfSymbols; j++ {
			finalString += symbolToPrint
		}
		finalString += "\n"
	}
	return finalString
}

// The goal of this function is to print a top-down triangle in the center
//with spaces between characters.
// As input to the function are numberOfRows and symbolToPrint.
func TriangleCenterTopDown(numberOfRows int, symbolToPrint string) string {
	var finalString = ""
	var numOfSpaces = ""

	for i := numberOfRows; i >= 1; i-- {
		numOfSpaces += " "
		var numberOfSymbols = 1 * i

		finalString += numOfSpaces
		for j := 0; j < numberOfSymbols; j++ {
			finalString += symbolToPrint + " "
		}
		finalString += "\n"
	}
	return finalString
}

// The goal of this function is to print triangle in the center
////with spaces between characters.
// As input to the function are numberOfRows and symbolToPrint.
func TriangleCenter(numberOfRows int, symbolToPrint string) string {
	var finalString = ""
	var numOfSpaces = ""

	for i := 0; i < numberOfRows; i++ {
		numOfSpaces = addSpacesToString(numberOfRows - i)
		var numberOfSymbols = 1 * (i + 1)

		finalString += numOfSpaces
		for j := 0; j < numberOfSymbols; j++ {
			finalString += symbolToPrint + " "
		}
		finalString += "\n"
	}
	return finalString
}

func addSpacesToString(numberOfSpaces int) string {
	var finalString = ""

	for i := 0; i < numberOfSpaces; i++ {
		finalString += " "
	}

	return finalString
}
