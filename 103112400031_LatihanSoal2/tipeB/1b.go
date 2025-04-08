// Savila Nur Fadilla
// 103112400031

package main

import "fmt"

func cekGanjil(a, b int) int {
	jumlah := 0
	for i := a; i <= b; i++ {
		if i % 2 != 0 {
			jumlah++
		}
	}
	return jumlah
}

func main() {
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	fmt.Print("Banyaknya angka ganjil: ", cekGanjil(a, b))
}