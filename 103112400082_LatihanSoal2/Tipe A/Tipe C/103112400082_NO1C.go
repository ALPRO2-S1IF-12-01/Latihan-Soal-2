// Rizkina Azizah_103112400082
package main

import (
	"fmt"
	"strconv"
)
 
func potongBilangan(n int) (int, int) {
	inputStr := strconv.Itoa(n)
	length := len(inputStr)

	mid := (length + 1) / 2 
	
	bilangan1, _ := strconv.Atoi(inputStr[:mid])
	bilangan2, _ := strconv.Atoi(inputStr[mid:])

	return bilangan1, bilangan2
}

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&n)

	bilangan1, bilangan2 := potongBilangan(n)

	fmt.Printf("Bilangan 1: %d\n", bilangan1)
	fmt.Printf("Bilangan 2: %d\n", bilangan2)
	fmt.Printf("Hasil penjumlahan: %d\n", bilangan1+bilangan2)
}