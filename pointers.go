package main

// Operator to create a pointer
// Operator to get the memory address &

import (
	"fmt"
)

func pointer() {
	//var ptr *int // Create a pointer with *
	//fmt.Println(prt)
}

type Node struct {
	prev int
	next int
	key  interface{}
}

// \* is used to create a pointer - reference
// & is used to de-reference
func main() {
	// n := 34
	// var a_pointer *int = &n
	//m := n
	
	// In simple words we can get the memory address from a variable
	// and pointing this memory address in other variable
	// in more simple words the two variables have the same value and
	// share the same memory address
	// fmt.Println(a_pointer)
	// fmt.Println(&n)
	var v int = 19
	var p *int = &v

	fmt.Printf("Las variable v es: %d \n", v)
	fmt.Printf("La direccion de memoria de v es: %v \n", &v)
	fmt.Printf("El puntero p hace referencia a la direccion de memoria: %v \n", p)
	fmt.Printf("Al desreferenciar el puntero p obtengo el valor: %d \n", *p)
}