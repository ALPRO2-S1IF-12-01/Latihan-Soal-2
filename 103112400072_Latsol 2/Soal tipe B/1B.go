package main

import "fmt"

func main() {

	var a, b int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	if a > b {

		a, b = b, a
	}

	jumlahGanjil := 0

	for i := a; i <= b; i++ {
		if i%2 != 0 {
			jumlahGanjil++
		}
	}

	fmt.Printf("Banyaknya angka ganjil: %d\n", jumlahGanjil)
}
