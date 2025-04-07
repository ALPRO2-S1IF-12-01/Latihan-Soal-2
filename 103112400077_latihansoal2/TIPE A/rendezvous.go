package main

import "fmt"

func pertemuan(x, y int) int {
	jumlah := 0
	for hari := 1; hari <= 365; hari++ {
		if hari%x == 0 && hari%y != 0 {
			jumlah++
		}
	}
	return jumlah
}

func main() {
	var x_103112400077, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x_103112400077)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	hasil := pertemuan(x_103112400077, y)
	fmt.Printf("Jumlah pertemuan dalam setahun: %d\n", hasil)
}
