// JESIKA METANIA RAHMA ARIFIN
// 103112400080

package main

import "fmt"

func hitungBiaya(jumlahMenu int) int {
	if jumlahMenu <= 3 {
		return 10000
	} else if jumlahMenu > 3 && jumlahMenu <= 50 {
		return 10000 + (jumlahMenu-3)*2500
	} else {
		return 100000
	}
}

func main() {
	var M int
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&M)

	for i := 1; i <= M; i++ {
		var jumlahMenu, banyakOrang, sisa int
		fmt.Print("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya)\n: ")
		fmt.Scan(&jumlahMenu, &banyakOrang, &sisa)

		totalBiaya := hitungBiaya(jumlahMenu)
		if sisa == 1 {
			totalBiaya *= banyakOrang
		}
		fmt.Printf("Total biaya untuk rombongan %d: Rp %d\n", i, totalBiaya)
	}
}