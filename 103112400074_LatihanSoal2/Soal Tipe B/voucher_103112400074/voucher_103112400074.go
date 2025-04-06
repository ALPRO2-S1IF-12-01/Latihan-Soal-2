//Ahmad Ruba'i
//103112400074
package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	hasil :=voucher(a, b)

	fmt.Println("Banyaknya angka ganjil:", hasil)
}

func voucher(a, b int) int {
	jumlah := 0
	for i := a; i <= b; i++ {
		if i%2 != 0 {
			jumlah++
		}
	}
	return jumlah
}
