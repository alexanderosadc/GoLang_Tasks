package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	var sortingWaitGroup sync.WaitGroup
	numberOfSlices := 4
	arr := *UserInput()
	//arr := []int{11, 22, 33, 44}
	resultArr, err := SeparateArray(arr, numberOfSlices)
	if err != nil {
		fmt.Println("ERROR:" + err.Error())
		return
	}

	sortingWaitGroup.Add(numberOfSlices)

	for _, element := range *resultArr {
		go BubbleSort(element, &sortingWaitGroup)
	}

	sortingWaitGroup.Wait()

	mergeSortedArrays(resultArr)
}

func mergeSortedArrays(arr [][]int) {
	finalArr := []int{}
	maximalArr := 0
}

func SeparateArray(arr []int, numberOfSlices int) (*[][]int, error) {
	if len(arr) < numberOfSlices {
		return nil, errors.New("Number of slices is higher than array length")
	}

	arrOfElements := [][]int{}
	arrOfIndexes := getStartEndIndexesFromArray(len(arr), numberOfSlices)

	for _, element := range *arrOfIndexes {
		startIndex := element[0]
		endIndex := element[1]
		arrOfElements = append(arrOfElements, arr[startIndex:endIndex])
	}

	return &arrOfElements, nil
}

func UserInput() *[]int {
	arr := []int{}
	reader := bufio.NewReader(os.Stdin)
	var command string
	var err error

	for {
		fmt.Println("Introduce integer number")
		fmt.Print(">")

		command, err = reader.ReadString('\n')
		commandWithoutEscChars := strings.TrimSpace(command)
		if err != nil {
			fmt.Println("ERROR: " + err.Error())
			continue
		} else if commandWithoutEscChars == "end" {
			break
		}

		integerElement, err := strconv.Atoi(commandWithoutEscChars)
		if err != nil {
			fmt.Println("ERROR: " + err.Error())
			continue
		}

		arr = append(arr, integerElement)
	}

	return &arr
}

func getStartEndIndexesFromArray(arrayLength int, numberOfSlices int) *[][]int {
	floatArrLength := float64(arrayLength)
	floatNumberOfSlices := float64(numberOfSlices)

	nrOfElementsInSlice := math.Floor(floatArrLength / floatNumberOfSlices)
	nrOfOverflowSlices := math.Mod(floatArrLength, floatNumberOfSlices)
	arrOfStartEndIndexPairs := [][]int{}

	for startIndex := 0; startIndex < arrayLength; {
		endIndex := startIndex + int(nrOfElementsInSlice)

		if int(nrOfOverflowSlices) > 0 {
			arrOfStartEndIndexPairs = append(arrOfStartEndIndexPairs, []int{startIndex, endIndex + 1})
			startIndex += int(nrOfElementsInSlice) + 1
			nrOfOverflowSlices--
		} else {
			arrOfStartEndIndexPairs = append(arrOfStartEndIndexPairs, []int{startIndex, endIndex})
			startIndex += int(nrOfElementsInSlice)
		}
	}

	return &arrOfStartEndIndexPairs
}

func BubbleSort(sliceOfNumbers []int, waitGroup *sync.WaitGroup) {
	fmt.Println("SubArray before sort: ", sliceOfNumbers)
	lengthOfSlice := len(sliceOfNumbers)

	for i := 0; i < lengthOfSlice-1; i++ {
		for j := 0; j < lengthOfSlice-i-1; j++ {

			if sliceOfNumbers[j] > sliceOfNumbers[j+1] {
				swap(&sliceOfNumbers[j], &sliceOfNumbers[j+1])
			}
		}
	}
	waitGroup.Done()
}

func swap(firstValue *int, secondValue *int) {
	numberBuffer := *firstValue
	*firstValue = *secondValue
	*secondValue = numberBuffer
}
