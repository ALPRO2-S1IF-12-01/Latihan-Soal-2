// Rizkina Azizah_103112400082
package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	pertemuan := 0
	for i := 1; i <= 365; i++ {
		if i%x == 0 && i%y != 0 {
			pertemuan++
		}
	}

	
	fmt.Printf("Jumlah pertemuan dalam setahun: %d\n", pertemuan)
}