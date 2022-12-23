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

func fibonacci_3(limit int) []int {
	initValues := []int {0, 1}

	for initValues[len(initValues) -1] < limit {
		initValues = append(initValues, (initValues[len(initValues) - 2] + initValues[len(initValues) - 1]))
	}

	return initValues
}

func main() {
	input := 300
	initValues := []int {0, 1}
	values := fibonacci(input, initValues, 0, 1)
	fmt.Println(values)

	values_2 := fibonacci_2(input, initValues)
	fmt.Println(values_2)

	values_3 := fibonacci_3(input)
	fmt.Println(values_3)
}