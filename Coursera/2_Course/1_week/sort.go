package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{}
	arr = UserInput(arr)
	BubbleSort(arr)
	fmt.Println(arr)
}

func UserInput(arr []int) []int {
	for len(arr) < 10 {
		var command string

		fmt.Println("Write number or stop app by entering X sign")
		fmt.Scan(&command)

		if command == "X" {
			break
		}

		enteredNumber, err := strconv.Atoi(command)
		if err != nil {
			fmt.Println("ERROR: " + err.Error())
			continue
		}
		// When appended new element new location in memory is created,
		// so it is not points to the same array and we need to return
		// new array.
		arr = append(arr, enteredNumber)
	}
	return arr
}

func BubbleSort(sliceOfNumbers []int) {
	lengthOfSlice := len(sliceOfNumbers)

	for i := 0; i < lengthOfSlice-1; i++ {
		for j := 0; j < lengthOfSlice-i-1; j++ {

			if sliceOfNumbers[j] > sliceOfNumbers[j+1] {
				Swap(&sliceOfNumbers[j], &sliceOfNumbers[j+1])
			}
		}
	}
}

func Swap(firstValue *int, secondValue *int) {
	numberBuffer := *firstValue
	*firstValue = *secondValue
	*secondValue = numberBuffer
}
