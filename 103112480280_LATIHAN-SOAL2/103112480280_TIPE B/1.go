// Nama: Anggun Wahyu Widiyana
// NIM : 103112480280
package main

import "fmt"

func main() {
	var a_103112480280, b int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a_103112480280)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	countGanjil := 0
	for i := a_103112480280; i <= b; i++ {
		if i%2 != 0 {
			countGanjil++
		}
	}

	fmt.Printf("Banyaknya angka ganjil: %d\n", countGanjil)
}