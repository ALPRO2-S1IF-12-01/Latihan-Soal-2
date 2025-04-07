// Rizkina Azizah_103112400082
package main

import "fmt"

func hitungBiaya(jumlahMenu, jumlahOrang, sisaMakanan int) int {
	var totalBiaya int

	if jumlahMenu > 50 {
		totalBiaya = 100000 
	} else if jumlahMenu <= 3 {
		totalBiaya = 10000 
	} else {
		totalBiaya = 10000 + (jumlahMenu-3)*2500 
	}

	if sisaMakanan == 1 {
		totalBiaya *= jumlahOrang
	}

	return totalBiaya
}

func main() {
	var M int

	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&M)

	for i := 1; i <= M; i++ {
		var jumlahMenu, jumlahOrang int
		var sisaMakanan int

		fmt.Printf("Masukkan jumlah menu, jumlah orang, dan status makanan yang tersisa (0 untuk tidak, 1 untuk iya): ")
		fmt.Scan(&jumlahMenu, &jumlahOrang, &sisaMakanan)

		totalBiaya := hitungBiaya(jumlahMenu, jumlahOrang, sisaMakanan)

		fmt.Printf("Total biaya rombongan %d: Rp %d\n", i, totalBiaya)
	}
}