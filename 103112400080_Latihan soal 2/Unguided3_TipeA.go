// JESIKA METANIA RAHMA ARIFIN
// 103112400080

package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	pertemuan := hitungPertemuan(x, y)
	fmt.Printf("Jumlah pertemuan dalam setahun: %d\n", pertemuan)
}

func hitungPertemuan(x, y int) int {
	jumlahHari := 0
	for hari := 1; hari <= 365; hari++ {
		if hari%x == 0 && hari%y != 0 {
			jumlahHari++
		}
	}

	return jumlahHari
}