package main

import "fmt"

func main() {
	var rombongan_103112400055 int
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&rombongan_103112400055)

	for i := 0; i < rombongan_103112400055; i++ {
		var menu, orang int
		var sisa bool

		fmt.Printf("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya): ")
		fmt.Scan(&menu, &orang, &sisa)

		totalBiaya := 10000
		if menu > 3 {
			totalBiaya += (menu - 3) * 2500
		}
		if menu > 50 {
			totalBiaya = 100000
		}

		if sisa {
			totalBiaya *= orang
		}

		fmt.Printf("Total biaya untuk rombongan %d: Rp %d\n", i+1, totalBiaya)
	}
}
