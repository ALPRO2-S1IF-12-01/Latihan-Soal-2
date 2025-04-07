// JESIKA METANIA RAHMA ARIFIN
// 103112400080

package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	if a > b {
		fmt.Print("Nilai a harus lebih kecil atau sama dengan b: ")
		return
	}

	count := 0
	for i := a; i <= b; i++ {
		if i%2 != 0 {
			count++
	}
	}
	fmt.Printf("Banyaknya angka ganjil dari %d hingga %d: %d\n", a, b, count )
}