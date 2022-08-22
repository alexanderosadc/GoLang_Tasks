package main

import (
	"fmt"
	"strings"
)

func main() {
	var userStr string

	fmt.Printf("Write text: \n")
	// All input should be placed in quotes because if we will have spaces they will
	//be stored as a space separated massive
	_, err := fmt.Scan(&userStr)

	if err != nil {
		fmt.Println("ERROR: " + err.Error() + "\n")
	}

	// Transforms string in lowercase letters
	var lowerUserStr = strings.ToLower(userStr)
	fmt.Println("Lowered case string: " + userStr)
	if strings.HasPrefix(lowerUserStr, "i") &&
		strings.HasSuffix(lowerUserStr, "n") &&
		strings.Contains(lowerUserStr, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
