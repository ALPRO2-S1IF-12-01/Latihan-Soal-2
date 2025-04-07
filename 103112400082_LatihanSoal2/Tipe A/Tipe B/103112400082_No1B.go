// Rizkina Azizah_103112400082
package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	if a > b {
		fmt.Println("Nilai a harus kurang dari atau sama dengan b.")
		return
	}

	jumlahGanjil := 0
	for i := a; i <= b; i++ {
		if i%2 != 0 {
			jumlahGanjil++
		}
	}

	fmt.Print("Banyaknya angka ganjil : ", jumlahGanjil)
}