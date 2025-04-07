// RYAN AKEYLA NOVIANTO WIDODO
// 103112400081

package main

import "fmt"

func hitungPertemuan(x, y int) int {
	jumlahPertemuan := 0
	for hari := 1; hari <= 365; hari++ {
		if hari%x == 0 && hari%y != 0 {
			jumlahPertemuan++
		}
	}
	return jumlahPertemuan
}

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x:")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y:")
	fmt.Scan(&y)

	jumlahPertemuan := hitungPertemuan(x, y)
	fmt.Printf("Jumlah Pertemuan dalam setahun: %d\n", jumlahPertemuan)
}
