// Savila Nur Fadilla
// 103112400031

package main

import "fmt"

func hitungPertemuan(x, y int) int {
	jumlahPertemuan := 0
	for i := 1; i <= 365; i++ {
		if i % x == 0 && i % y != 0 {
			jumlahPertemuan++
		}
	}
	return jumlahPertemuan
}

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	fmt.Printf("Jumlah pertemuan dalam setahun: %v", hitungPertemuan(x, y))
}