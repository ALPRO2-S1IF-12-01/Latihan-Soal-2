// RYAN AKEYLA NOVIANTO WIDODO
// 103112400081

package main

import "fmt"

func main() {
	var a, b_103112400081 int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b_103112400081)

	if a > b_103112400081 {
		fmt.Println("Kesalahan: Nilai a harus lebih kecil atau sama dengan b.")
		return
	}

	count := 0
	for i := a; i <= b_103112400081; i++ {
		if i%2 != 0 {
			count++
		}
	}

	fmt.Printf("Banyaknya angka ganjil: %d", count)
}
