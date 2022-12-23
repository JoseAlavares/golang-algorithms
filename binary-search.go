package main

import (
	"fmt"
	"sort"
)

func binary_search(input_value int, arr []int, min int, max int) (bool, int) {
	sort.Ints(arr)

	if min > max {
		return false, 0
	}

	var middle int = (min + max) / 2

	if input_value == arr[middle] {
		return true, middle
	} else if input_value < arr[middle] {
		return binary_search(input_value, arr, min, middle - 1)
	} else if input_value > arr[middle] {
		return binary_search(input_value, arr, middle + 1, max)
	}

	return false, -1
}

func main() {
	unordered_values := []int {0, 2, 10, 30, 24, 47, 100, 11, 37, 19, 34, 45, 78, 102, 23}
	exists, index := binary_search(10, unordered_values, 0, len(unordered_values) - 1)

	if exists {
		fmt.Println("Exists the value in the position:", index + 1)
	} else {
		fmt.Println("Not exists the value")
	}
}