package main

import "fmt"

const base = 10 // 0 - 9 digits

func radixSort(array []int) []int {
	if len(array) == 0 {
		return array
	}

	maxValue := max(array)

	placeValue := 1
	for maxValue/placeValue > 0 {
		countingSort(array, placeValue)
		fmt.Println("place value: ", placeValue, "array: ", array)
		placeValue *= base
	}
	return array
}

// use counting sort as subroutine
func countingSort(array []int, placeValue int) {
	minValue := getMinValue(array, placeValue)
	offset := 0 // deal with negative values
	if minValue < 0 {
		offset = 0 - minValue
	}
	counts := make([]int, base+offset) // minValue-9 digits
	sortedArray := make([]int, len(array))

	for i := range array {
		countsIdx := getDigit(array, i, placeValue) + offset
		counts[countsIdx] += 1
	}

	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	for i := len(array) - 1; i >= 0; i-- {
		countsIdx := getDigit(array, i, placeValue) + offset
		position := counts[countsIdx]
		sortedIdx := position - 1
		sortedArray[sortedIdx] = array[i]
		counts[countsIdx] -= 1
	}

	// copy values to original array
	for i := range array {
		array[i] = sortedArray[i]
	}
}

// get minimum value of each digit group
func getMinValue(array []int, placeValue int) int {
	minValue := getDigit(array, 0, placeValue)
	for i := 1; i < len(array); i++ {
		element := getDigit(array, i, placeValue)
		if minValue > element {
			minValue = element
		}
	}
	return minValue
}

// ex. number: 2022, place value: 100 => digit = 0
// ex. number: -1223, place value: 100 => digit = -2
func getDigit(array []int, idx int, placeValue int) int {
	element := array[idx]
	digit := abs(element) / placeValue % base
	if element < 0 {
		return -digit
	}
	return digit
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func max(array []int) int {
	maxValue := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > maxValue {
			maxValue = array[i]
		}
	}
	return maxValue
}

func main() {
	arrays := [][]int{
		{2022, 5, 398, 159, 16, 1223, 7},
		{2022, -5, 398, 159, 16, -1223, 7},
	}

	for _, array := range arrays {
		fmt.Println("unsorted array: ", array)
		fmt.Println("sorted array: ", radixSort(array))
		fmt.Println()
	}
}

/* output:
unsorted array:  [2022 5 398 159 16 1223 7]
place value:  1 array:  [2022 1223 5 16 7 398 159]
place value:  10 array:  [5 7 16 2022 1223 159 398]
place value:  100 array:  [5 7 16 2022 159 1223 398]
place value:  1000 array:  [5 7 16 159 398 1223 2022]
sorted array:  [5 7 16 159 398 1223 2022]

unsorted array:  [2022 -5 398 159 16 -1223 7]
place value:  1 array:  [-5 -1223 2022 16 7 398 159]
place value:  10 array:  [-1223 -5 7 16 2022 159 398]
place value:  100 array:  [-1223 -5 7 16 2022 159 398]
place value:  1000 array:  [-1223 -5 7 16 159 398 2022]
sorted array:  [-1223 -5 7 16 159 398 2022]
*/
