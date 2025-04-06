//Anastasia Adinda N.I
package main

import (
	"fmt"
)

func main() {
	var a, b int
	var jumlahGanjil_103112400085 int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	if a > b {
		fmt.Println("Nilai a harus lebih kecil atau sama dengan b.")
		return
	}

	for i := a; i <= b; i++ {
		if i%2 != 0 {
			jumlahGanjil_103112400085++
		}
	}

	fmt.Printf("Banyaknya angka ganjil: %d\n", jumlahGanjil_103112400085)
}
