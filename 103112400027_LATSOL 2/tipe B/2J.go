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
	fmt.Println("Nama: Raja Muhammad Lufhti\nNim : 103112400027")
	var M int

	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&M)

	for i := 1; i <= M; i++ {
		var jumlahMenu, jumlahOrang, sisaMakanan int

		fmt.Printf("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya) untuk rombongan %d: ", i)
		fmt.Scan(&jumlahMenu, &jumlahOrang, &sisaMakanan)

		biaya := hitungBiaya(jumlahMenu)

		if sisaMakanan == 1 {
			biaya *= jumlahOrang
		}

		fmt.Printf("Total biaya untuk rombongan %d: Rp %d\n", i, biaya)
	}
}
