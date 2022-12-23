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

func main() {
	exists_value := linear_search(31)

	if exists_value {
		fmt.Println("Exists the value")
	} else {
		fmt.Println("Not exists the value")
	}
}

