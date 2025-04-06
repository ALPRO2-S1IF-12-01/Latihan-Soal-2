// Felix Pedrosa V

package main

import "fmt"

func hitungBiaya(jumlahMenu int) int {
	if jumlahMenu <= 3 {
		return 10000
	} else if jumlahMenu > 50 {
		return 100000
	} else {
		return 10000 + (jumlahMenu-3)*2500
	}
}

func main() {
	var rombongan_103112400056 int
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&rombongan_103112400056)

	totalBiaya := 0

	for i := 1; i <= rombongan_103112400056; i++ {
		var jumlahMenu, jumlahOrang, sisaMakanan int

		fmt.Printf("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya) %d: ", i)
		fmt.Scan(&jumlahMenu, &jumlahOrang, &sisaMakanan)

		biaya := hitungBiaya(jumlahMenu)

		if sisaMakanan == 1 {
			biaya *= jumlahOrang
		}

		totalBiaya += biaya
		fmt.Printf("Total biaya untuk rombongan %d: Rp %d\n", i, biaya)
	} 
}
