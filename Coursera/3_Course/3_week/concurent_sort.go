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
	bubblesortSliceOfSlices(*resultArr)
	fmt.Println(*mergeArrays(*resultArr))
}

// Iterates over slice of slices and applies function for merging all slice.
// Returns one slice with all elements from array sorted
func mergeArrays(arr [][]int) *[]int {
	finalSlice := []int{}

	finalSlice = append(finalSlice, arr[0]...)

	for index := 1; index < len(arr); index++ {
		finalSlice = *merge2Arrays(finalSlice, arr[index])
	}

	return &finalSlice
}

// Alghorithm for merging 2 slices in the ascending order of elements.
// arr1 and arr2 should be sorted arrays
func merge2Arrays(arr1 []int, arr2 []int) *[]int {
	arr3 := []int{}
	i := 0
	j := 0
	n1 := len(arr1)
	n2 := len(arr2)

	for i < n1 && j < n2 {
		if arr1[i] < arr2[j] {
			arr3 = append(arr3, arr1[i])
			i++
		} else {
			arr3 = append(arr3, arr2[j])
			j++
		}
	}

	for i < n1 {
		arr3 = append(arr3, arr1[i])
		i++
	}

	for j < n2 {
		arr3 = append(arr3, arr2[j])
		j++
	}

	return &arr3
}

// Sortes slice of slices based on the first element of each slice
func bubblesortSliceOfSlices(arr [][]int) {
	lengthOfSlice := len(arr)

	for i := 0; i < lengthOfSlice-1; i++ {
		for j := 0; j < lengthOfSlice-i-1; j++ {

			if arr[j][0] > arr[j+1][0] {

				arrBuffer := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = arrBuffer
			}
		}
	}
}

// Separates slice in number of slices.
// Returns slice of slices
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

// Reads numbers from the console, parses them into int and returns array.
// When user wants to stop introducing numbers he introduces end word to the console.
func UserInput() *[]int {
	arr := []int{}
	reader := bufio.NewReader(os.Stdin)
	var command string
	var err error

	for {
		fmt.Println("Introduce integer number or \"end\" command to stop introducing numbers")
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

// Extracts start and end index for separating one slice into many.
// returns slice of slices. In them there are pair of start and end index.
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
