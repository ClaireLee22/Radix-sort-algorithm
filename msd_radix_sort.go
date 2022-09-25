package main

import "fmt"

const base = 10

func radixSort(array []int) []int {

	maxValue := max(array)
	// leftmost place value
	placeValue := pow(base, getNumberOfDigits(maxValue)-1)

	bucketSort(array, placeValue)

	return array
}

func bucketSort(array []int, placeValue int) []int {
	if placeValue < 1 {
		return array
	}

	if len(array) <= 1 {
		return array
	}

	buckets := [][]int{}
	minValue := getMinValue(array, placeValue)

	offset := 0
	if minValue < 0 {
		offset = 0 - minValue
	}
	for i := 0; i < (base + offset); i++ {
		buckets = append(buckets, []int{})
	}

	for idx, num := range array {
		bucketIdx := getDigit(array, idx, placeValue) + offset
		buckets[bucketIdx] = append(buckets[bucketIdx], num)
	}
	fmt.Println("place value: ", placeValue, "bucket: ", buckets)

	sortedIdx := 0
	for _, bucket := range buckets {
		sortedBucket := bucketSort(bucket, placeValue/base)
		for _, num := range sortedBucket {
			array[sortedIdx] = num
			sortedIdx += 1
		}
	}
	fmt.Println("place value: ", placeValue, "array: ", array)

	return array
}

func getNumberOfDigits(num int) int {
	numberOfDigits := 0
	for num > 0 {
		numberOfDigits += 1
		num /= base
	}
	return numberOfDigits
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

func pow(a int, power int) int {
	result := 1
	for i := 0; i < power; i++ {
		result *= a
	}
	return result
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

/* outout
unsorted array:  [2022 5 398 159 16 1223 7]
place value:  1000 bucket:  [[5 398 159 16 7] [1223] [2022] [] [] [] [] [] [] []]
place value:  100 bucket:  [[5 16 7] [159] [] [398] [] [] [] [] [] []]
place value:  10 bucket:  [[5 7] [16] [] [] [] [] [] [] [] []]
place value:  1 bucket:  [[] [] [] [] [] [5] [] [7] [] []]
place value:  1 array:  [5 7]
place value:  10 array:  [5 7 16]
place value:  100 array:  [5 7 16 159 398]
place value:  1000 array:  [5 7 16 159 398 1223 2022]
sorted array:  [5 7 16 159 398 1223 2022]

unsorted array:  [2022 -5 398 159 16 -1223 7]
place value:  1000 bucket:  [[-1223] [-5 398 159 16 7] [] [2022] [] [] [] [] [] [] []]
place value:  100 bucket:  [[-5 16 7] [159] [] [398] [] [] [] [] [] []]
place value:  10 bucket:  [[-5 7] [16] [] [] [] [] [] [] [] []]
place value:  1 bucket:  [[-5] [] [] [] [] [] [] [] [] [] [] [] [7] [] []]
place value:  1 array:  [-5 7]
place value:  10 array:  [-5 7 16]
place value:  100 array:  [-5 7 16 159 398]
place value:  1000 array:  [-1223 -5 7 16 159 398 2022]
sorted array:  [-1223 -5 7 16 159 398 2022]
*/
