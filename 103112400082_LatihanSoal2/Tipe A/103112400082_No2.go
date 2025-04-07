// Rizkina Azizah_103112400082
package main

import "fmt"

func PerfectNumber(n int) bool {
	cek := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			cek += i
		}
	}
	return cek == n
}

func main() {
	var a, b int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	fmt.Printf("Perfect number antara %d dan %d: ", a, b)
	for i := a; i <= b; i++ {
		if PerfectNumber(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println() 
}