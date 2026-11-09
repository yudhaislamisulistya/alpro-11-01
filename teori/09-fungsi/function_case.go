package main

import "fmt"

func hitung(x, y int) int { // x dan y adalah sebuah parameter
	hasil := 2*x + y - 5
	return hasil
}

func main() {
	a := 10
	fmt.Println(hitung(a, 5)) // a dan 5 ini adalah sebuah argumen
	fmt.Println(hitung(20, 10))
	fmt.Println(hitung(30, 1))
}
