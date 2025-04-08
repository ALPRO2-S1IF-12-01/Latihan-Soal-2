// Savila Nur Fadilla
// 103112400031

package main

import "fmt"

func hitungBiaya(menu, orang int, sisa bool) int {
	var biaya int

	if menu <= 3 {
		biaya = 10000
	} else if menu > 50 {
		biaya = 100000
	} else {
		biaya = 10000 + (menu - 3) * 2500
	}

	if sisa {
		biaya *= orang
	}
	return biaya
}

func main() {
	var M, menu, orang, sisaInt int
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&M)

	for i := 1; i <= M; i++ {
		fmt.Print("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya): ")
		fmt.Scan(&menu, &orang, &sisaInt)

		sisa := sisaInt == 1
		biaya := hitungBiaya(menu, orang, sisa)

		fmt.Printf("Total biaya untuk rombongan %v: Rp %v\n", i, biaya)
	}
}