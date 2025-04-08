package main

import (
	"fmt"
)

func main() {
	var x_103112400087, y int

	fmt.Print("masukkan nilai x: ")
	fmt.Scan(&x_103112400087)
	fmt.Print("masukkan nilai y: ")
	fmt.Scan(&y)

	jumlahpertemuan := 0

	for hari := 1; hari <= 365; hari++ {
		if hari%x_103112400087 == 0 && hari%y != 0 {
			jumlahpertemuan++
		}
	}

	fmt.Printf("jumlah pertemuan dalam setahun: %d\n", jumlahpertemuan)
}
