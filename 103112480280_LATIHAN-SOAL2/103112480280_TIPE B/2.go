// Nama: Anggun Wahyu Widiyana
// NIM : 103112480280
package main

import "fmt"

func main() {
	var M_103112480280 int
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&M_103112480280)

	var totalBayar int = 0

	for i := 1; i <= M_103112480280; i++ {
		var menu, orang int
		var sisa int
		fmt.Print("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0/1): ")
		fmt.Scan(&menu, &orang, &sisa)

		var biaya int
		if menu > 50 {
			biaya = 100000
		} else if menu > 3 {
			biaya = 10000 + (menu-3)*2500
		} else {
			biaya = 10000
		}

		if sisa == 1 {
			biaya = biaya * orang
		}

		fmt.Printf("Total biaya untuk rombongan %d: Rp %d\n", i, biaya)
		totalBayar += biaya
	}
}