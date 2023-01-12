package main

import (
	"fmt"
	"time"
	"sync"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main_() {
	// Example using Wait Groups
	var wg sync.WaitGroup

	fmt.Println("Hello")
	wg.Add(1)
	go say("World", &wg)
	wg.Wait()
	go func(text string) {
		fmt.Println(text)
	}("Adios")
	time.Sleep(time.Second * 1)
}

func say2(text string, c chan<- string) {
	c <- text
}
func main() {
	c := make(chan string, 1)

	fmt.Println("Hello")
	go say2("Bye", c)
	fmt.Println(<- c)
}