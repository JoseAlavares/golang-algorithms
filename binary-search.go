package main

import (
	"fmt"
	"math"
	"sort"
)

func binary_search(input_value int, index int, arr []) bool {
	length := 15
	division := float64(length/2)
	middle := int64(math.Floor(division))

	if index >= length - 1 {
		return false
	}

	if unordered_values[middle] <= input_value {
		unordered_values_slice := unordered_values[1:middle]

		if input_value == unordered_values_slice[index] {
			return true
		}

		return binary_search(input_value, index + 1)
	}

	if unordered_values[middle] >= input_value {
		unordered_values_slice := unordered_values[middle:length - 1]

		if input_value == unordered_values_slice[index] {
			return true
		}

		return binary_search(input_value, index + 1)
	}

	return false
}

func main() {
	unordered_values := []int{0, 2, 10, 30, 24, 47, 100, 11, 37, 19, 34, 45, 78, 102, 23}
	sort.Ints(unordered_values)
	exists := binary_search(10, 0, unordered_values)

	if exists {
		fmt.Println("Exists the value")
	} else {
		fmt.Println("Not exists the value")
	}
}