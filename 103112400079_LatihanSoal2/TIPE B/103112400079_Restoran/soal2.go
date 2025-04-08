package main

import "fmt"

func main() {
	var jumlahRombongan_103112400079 int
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&jumlahRombongan_103112400079)

	for i := 1; i <= jumlahRombongan_103112400079; i++ {
		var menu, orang, sisa int
		fmt.Print("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya): ")
		fmt.Scan(&menu, &orang, &sisa)

		var tarif int

		if menu <= 3 {
			tarif = 10000
		} else if menu <= 50 {
			tarif = 10000 + (menu-3)*2500
		} else {
			tarif = 100000
		}

		if sisa == 1 {
			tarif *= orang
		}

		fmt.Printf("Total biaya untuk rombongan %d : Rp %d\n", i, tarif)
	}
}
