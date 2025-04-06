package main

import (
	"fmt"
)
func isPerfectNumber(n int) bool {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}
func main() {
	var a, b int
	var NIM string = "103112400076" 
	fmt.Println("NIM :", NIM) 
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Printf("Perfect numbers antara %d dan %d: ", a, b)
	ditemukan := false
	for i := a; i <= b; i++ {
		if isPerfectNumber(i) {
			fmt.Printf("%d ", i)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Print("Tidak ada")
	}
	fmt.Println()
}
