package main

import (
	"fmt"
	// "errors"
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func sentinelSearch(arr []int, key int) (bool, error) {
	if len(arr) <= 0 {
		arr = []int{ 10, 20, 180, 40, 30, 220, 198, 10002, 35, 2782 }
	}

	var length = len(arr)
	var i int = 0
	var last int = arr[length - 1]
	arr[length - 1] = key

	for arr[i] != key {
		i++
	}

	arr[length - 1] = last

	if i < length - 1 || arr[i] == key {
		return true, nil
	}

	return false, nil
}

func validateNumbers(str string) (bool, error) {
	matched, err := regexp.MatchString(`[0-9]`, str)

	if err != nil {
		fmt.Println("Error in the function validateNumbers")
		return false, err
	}

	return matched, err
}

func main()  {
	fmt.Println("Please enter the a number to fill the array")
	scanner := bufio.NewScanner(os.Stdin)
	arr := []int{}
	limit := 10

	for i:=0; i<limit; i++ {
		scanner.Scan()
		number := scanner.Text()

		matched, _ := validateNumbers(number)

		if !matched {
			fmt.Println("Not is a number the value of the array will declared by default")
			break
		}

		_number, _ := strconv.Atoi(number)
		arr = append(arr, _number)
	}

	fmt.Println("Insert the value to search")
	scanner.Scan()
	key := scanner.Text()
	matched, _ := validateNumbers(key)

	if !matched {
		fmt.Println("The value not is a number")
		return
	}

	_key, _ := strconv.Atoi(key)
	exists, _ := sentinelSearch(arr, _key)

	if exists {
		fmt.Println("The value exists")
		return
	}

	fmt.Println("The value not exists")
	return
}