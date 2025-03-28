// Raihan Adi Arb

package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	fmt.Printf("Banyaknya angka ganjil: %d\n", hitungJumlahGanjil(a, b))
}

func hitungJumlahGanjil(min, max int) int {
	jumlah := 0
	for angka := min; angka <= max; angka++ {
		if isGanjil(angka) {
			jumlah++
		}
	}
	return jumlah
}

func isGanjil(n int) bool {
	return n%2 != 0
}
