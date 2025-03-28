package main

import "fmt"

func main() {
	var jumlahRombongan int

	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&jumlahRombongan)

	for i := 1; i <= jumlahRombongan; i++ {
		var menu, orang, sisa int
		fmt.Print("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya): ")
		fmt.Scan(&menu, &orang, &sisa)

		totalBiaya := hitungTotalBiaya(menu, orang, sisa)
		fmt.Printf("Total biaya untuk rombongan %d: Rp %d\n", i, totalBiaya)
	}
}

func hitungTotalBiaya(menu, orang, sisa int) int {
	biaya := biayaMenu(menu)
	if sisa == 1 {
		return biaya * orang
	}
	return biaya
}

func biayaMenu(menu int) int {
	if menu > 50 {
		return 100000
	} else if menu <= 3 {
		return 10000
	}
	// Jika menu > 3 dan <= 50
	return 10000 + (menu-3)*2500
}
