package main

import "fmt"

func main() {
	var x_103112400055, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x_103112400055)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	jumlahPertemuan := 0
	for hari := 1; hari <= 365; hari++ {
		if hari%x_103112400055 == 0 && hari%y != 0 {
			jumlahPertemuan++
		}
	}

	fmt.Printf("Jumlah pertemuan dalam setahun: %d\n", jumlahPertemuan)
}
