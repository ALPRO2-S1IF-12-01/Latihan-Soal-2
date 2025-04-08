package main

import "fmt"

func main() {
	var a, b_103112400087 int
	fmt.Print("masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("masukkan nilai b: ")
	fmt.Scan(&b_103112400087)

	count := hitungGanjil(a, b_103112400087)
	fmt.Printf("banyaknya angka ganjil: %d\n", count)
}

func hitungGanjil(a, b_103112400087 int) int {
	count := 0
	for i := a; i <= b_103112400087; i++ {
		if i%2 != 0 {
			count++
		}
	}
	return count
}
