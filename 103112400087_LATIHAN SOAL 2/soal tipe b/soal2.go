package main

import "fmt"

func main() {
	var jumlahRombongan int
	fmt.Print("masukkan jumlah rombongan: ")
	fmt.Scan(&jumlahRombongan)

	for i := 1; i <= jumlahRombongan; i++ {
		var menu_103112400087, orang, sisa int
		fmt.Printf("masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya): ")
		fmt.Scan(&menu_103112400087, &orang, &sisa)

		biaya := hitungBiaya(menu_103112400087, orang, sisa)
		fmt.Printf("total biaya untuk rombongan %d: Rp %d\n", i, biaya)
	}
}

func hitungBiaya(menu_103112400087, orang, sisa int) int {
	biaya := 0
	if menu_103112400087 > 50 {
		biaya = 100000
	} else {
		if menu_103112400087 <= 3 {
			biaya = 10000
		} else {
			biaya = 10000 + (menu_103112400087-3)*2500
		}
	}

	if sisa == 1 {
		biaya *= orang
	}

	return biaya
}
