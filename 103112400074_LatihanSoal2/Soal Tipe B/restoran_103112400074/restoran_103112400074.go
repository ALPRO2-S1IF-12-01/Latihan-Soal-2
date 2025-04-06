//Ahmad Ruba'i
//103112400074
package main

import "fmt"

func main() {
	var jumlahRombongan int
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&jumlahRombongan)
	TotalBiaya(jumlahRombongan)
}

func TotalBiaya(jumlah int) {
	var menu, orang, Sisa int
	for i := 1; i <= jumlah; i++ {
		fmt.Println("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya): ")
		fmt.Scan(&menu, &orang, &Sisa)
		biaya := hitungHarga(menu)
		if Sisa == 1 {
			biaya *= orang
		}
		fmt.Println("Total biaya untuk rombongan", i, ": Rp", biaya)
	}
}
func hitungHarga(menu int) int {
	if menu > 50 {
		return 100000
	} else if menu > 3 {
		extra := menu - 3
		return 10000 + (extra * 2500)
	}
	return 10000
}
