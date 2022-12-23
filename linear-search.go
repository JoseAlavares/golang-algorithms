package main

import "fmt"

func linear_search(input_search int) bool {
	elements := [10]int{0, 2, 10, 30, 24, 47, 100, 11, 37, 19}

	for i:=0;i<10;i++ {
		if elements[i] == input_search {
			return true
		}
	}

	return false
}

func linear_search_recursive(arr []int, input_search int, start int, end int) bool {
	if start > end {
		return false
	}

	if input_search == arr[start] {
		return true
	} else {
		return linear_search_recursive(arr, input_search, start + 1, end)
	}

}

func main() {
	value_to_search := 11
	exists_value := linear_search(value_to_search)

	if exists_value {
		fmt.Println("Exists the value")
	} else {
		fmt.Println("Not exists the value")
	}

	elements := []int {0, 2, 10, 30, 24, 47, 100, 11, 37, 19}
	start := 0
	end := len(elements) - 1
	exists_value_recursive := linear_search_recursive(elements, value_to_search, start, end)

	if exists_value_recursive {
		fmt.Println("Exists the value in the recursive function")
	} else {
		fmt.Println("Not exists the value recursive")
	}
}

