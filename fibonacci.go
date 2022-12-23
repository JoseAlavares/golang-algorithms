package main

import "fmt"

func fibonacci(limit int, values []int, start int, end int) []int {
	if values[len(values) - 1] >= limit {
		return values
	}

	nextValue := values[start] + values[end]
	values = append(values, nextValue)

	return fibonacci(limit, values, start + 1, end + 1)
}

func fibonacci_2(limit int, values []int) []int {
	if values[len(values) - 1] >= limit {
		return values
	}

	nextValue := values[len(values) - 2] +  values[len(values) - 1]
	values = append(values, nextValue)
	return fibonacci_2(limit, values)
}

func main() {
	input := 300
	initValues := []int {0, 1}
	values := fibonacci(input, initValues, 0, 1)
	fmt.Println(values)
	
	values_2 := fibonacci_2(input, initValues)
	fmt.Println(values_2)
}