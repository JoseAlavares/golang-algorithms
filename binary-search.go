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

func binary_search_reduce_array(arr []int, input_search int) (bool, int) {
	sort.Ints(arr)
	middle := len(arr) - 1

	if input_search == arr[middle] {
		return true, middle
	}

	if input_search > arr[middle] {
		return binary_search_reduce_array(arr[middle + 1:len(arr) - 1], input_search)
	} else if input_search < arr[middle] {
		return binary_search_reduce_array(arr[0:middle - 1], input_search)
	} else {
		return false, -1
	}
}

func main() {
	input_search := 10
	unordered_values := []int {0, 2, 10, 30, 24, 47, 100, 11, 37, 19, 34, 45, 78, 102, 23}
	exists, index := binary_search(input_search, unordered_values, 0, len(unordered_values) - 1)

	if exists {
		fmt.Println("Exists the value in the position:", index + 1)
	} else {
		fmt.Println("Not exists the value")
	}

	exists_recursive, index_recursive := binary_search_reduce_array(unordered_values, input_search)

	if exists_recursive {
		fmt.Println("Exists the value recursive in the position:", index_recursive + 1)
	} else {
		fmt.Println("Not exists the value recursive")
	}
}