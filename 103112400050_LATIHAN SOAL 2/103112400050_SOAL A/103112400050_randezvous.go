package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	fmt.Printf("Jumlah pertemuan rahasia dalam setahun: %d\n", hitungPertemuan(x, y))
}

func hitungPertemuan(x, y int) (count int) {
	for hari := 1; hari <= 365; hari++ {
		if hari%x == 0 && hari%y != 0 {
			count++
		}
	}
	return
}
