package main

import "fmt"

func recursiveSummation(values []int, total int) int {
	if len(values) == 0 {
		return total
	}

	total += values[0]
	return recursiveSummation(values[1:], total)
}

func main() {
	values := []int {1, 2, 3, 4, 5}
	total := 0
	_total := recursiveSummation(values, total)
	fmt.Println(_total)
}