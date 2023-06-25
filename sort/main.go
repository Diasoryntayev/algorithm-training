package main

import (
	"fmt"
)

// type sorts interface {
// 	bubbleSort([]int) ([]int, error)
//  selectionSort([]int) ([]int, error)
//  insertionSort([]int) ([]int, error)
// }

func main() {
	arr := []int{11, 20, 54, 95, 23, 123, 9, 0}
	sortArrByBubbleSort, err := bubbleSort(arr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sortArrByBubbleSort)

	fmt.Println(selectionSort(arr))
	fmt.Println(insertionSort(arr))
}

func bubbleSort(numbers []int) ([]int, error) {
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers, nil
}

func selectionSort(numbers []int) ([]int, error) {
	for i := 0; i < len(numbers)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(numbers); j++ {
			if numbers[minIndex] > numbers[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			numbers[minIndex], numbers[i] = numbers[i], numbers[minIndex]
		}
	}
	return numbers, nil
}

func insertionSort(numbers []int) ([]int, error) {
	for i := 1; i < len(numbers); i++ {
		sorted := i - 1
		for sorted != -1 && numbers[sorted] > numbers[sorted+1] {
			numbers[sorted], numbers[sorted+1] = numbers[sorted+1], numbers[sorted]
			sorted--
		}
	}
	return numbers, nil
}
