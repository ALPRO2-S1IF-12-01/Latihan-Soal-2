//Pratama Bintang Daniswara 103112400051
package main

import "fmt"

func main() {
	var m int
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&m)

	for i := 1; i <= m; i++ {
		var menu, org, sisa, biaya int
		fmt.Print("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya): ")
		fmt.Scan(&menu, &org, &sisa)

		switch {
		case menu <= 3:
			biaya = 10000
		case menu <= 50:
			biaya = 10000 + (menu-3)*2500
		default:
			biaya = 100000
		}

		if sisa == 1 {
			biaya = biaya * org
		}

		fmt.Println("Total biaya untuk rombongan", i, ": Rp", biaya)
	}
}
