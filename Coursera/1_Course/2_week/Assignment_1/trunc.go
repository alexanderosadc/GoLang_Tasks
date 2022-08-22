package main

import (
	"fmt"
	"strings"
)

func main() {
	// Used float32 because in this program we don't care about precision
	var floatNumber float32
	var command string

	for {
		fmt.Printf("Write floating point(real) number: \n")
		// Returns errors for each element which was is not a number enetring in float32
		_, err := fmt.Scan(&floatNumber)

		if err != nil {
			fmt.Println("ERROR: " + err.Error() + "\n")
			continue
		}

		var intNumber = int(floatNumber)
		fmt.Printf("Integer number is %d \n", intNumber)
		fmt.Printf("Do you want to exit program? (Write Yes or No to the console)\n")
		_, err = fmt.Scan(&command)

		if err != nil {
			fmt.Println("ERROR: " + err.Error() + "\n")
			continue
		}
		// Checks if the answer from user for exiting program was Yes. Stops infinite loop
		if strings.Contains(command, "Yes") {
			break
		}
	}
}
