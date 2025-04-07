package main

import "fmt"

func hitungBiaya(menu, orang, sisa int) int {
	biaya := 0

	if menu <= 3 {
		biaya = 10000
	} else if menu <= 50 {
		biaya = 10000 + (menu-3)*2500
	} else {
		biaya = 100000								
	}

	if sisa == 1 {
		biaya *= orang
	}

	return biaya
}

func main() {
	var M_103112400077 int
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&M_103112400077)

	for i := 1; i <= M_103112400077; i++ {
		var menu, orang, sisa int
		fmt.Printf("Masukkan jumlah menu, jumlah orang, dan status sisa makanan untuk rombongan %d: ", i)
		fmt.Scan(&menu, &orang, &sisa)

		totalBiaya := hitungBiaya(menu, orang, sisa)
		fmt.Printf("Total biaya untuk rombongan %d: Rp %d\n", i, totalBiaya)
	}
}
