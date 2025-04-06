//Anastasia Adinda N.I
package main

import (
	"fmt"
)

func main() {
	var M int
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&M)

	for i := 1; i <= M; i++ {
		var menu, orang, sisa int
		var totalBiaya_103112400085 int

		fmt.Print("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya): ")
		fmt.Scan(&menu, &orang, &sisa)

		if menu > 50 {
			totalBiaya_103112400085 = 100000
		} else if menu <= 3 {
			totalBiaya_103112400085 = 10000
		} else {
			totalBiaya_103112400085 = 10000 + (menu-3)*2500
		}

		if sisa == 1 {
			totalBiaya_103112400085 *= orang
		}

		fmt.Printf("Total biaya untuk rombongan %d: Rp %d\n", i, totalBiaya_103112400085)
	}
}
