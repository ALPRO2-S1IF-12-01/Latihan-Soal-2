package main

import "fmt"

func main() {
	var jumlahRombongan int

	fmt.Print("masukkan jumlah rombongan : ")
	fmt.Scan(&jumlahRombongan)

	totalBiaya(jumlahRombongan)
}

func totalBiaya(banyakRombongan int) {
	var menu, orangRombongan, sisa, harga int

	for urutan := 1; urutan <= banyakRombongan; urutan++ {
		fmt.Printf("\nrombongan %d:\n", urutan)
		fmt.Print("masukkan jumlah menu, orang, dan apakah ada sisa (1 = ya, 0 = tidak) : ")
		fmt.Scan(&menu, &orangRombongan, &sisa)

		//berdasarkan jumlah menu
		if menu > 50 {
			harga = 100000
		} else if menu > 3 {
			harga = 10000 + 2500*menu
		} else {
			harga = 10000
		}

		// kalo ada sisa makanan, biaya dikali jumlah orang
		if sisa == 1 {
			harga = harga * orangRombongan
		}

		fmt.Printf("total biaya untuk rombongan %d : Rp %d\n", urutan, harga)
		fmt.Print("Sheila Stephanie A [103112400086]")
	}
}
