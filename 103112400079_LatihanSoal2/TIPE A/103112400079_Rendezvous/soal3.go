package main

import "fmt"

func main() {
	var x, y_103112400079 int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y_103112400079)

	jumlahPertemuan := 0
	for hari := 1; hari <= 365; hari++ {
		if hari%x == 0 && hari%y_103112400079 != 0 {
			jumlahPertemuan++
		}
	}

	fmt.Printf("Jumlah pertemuan dalam setahun: %d\n", jumlahPertemuan)
}
