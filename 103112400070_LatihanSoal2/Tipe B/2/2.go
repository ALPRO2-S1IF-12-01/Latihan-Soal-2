// Achmad Zulvan Nur Hakim 103112400070
package main

import 	"fmt"

func totalBiaya(menu, orang, sisa int) int {
	var biaya int
	if menu <= 3 {
		biaya = 10000
	} else if menu <= 50 {
		biaya = 10000 + (menu-3)*2500
	} else {
		biaya = 100000
	}
	if sisa == 1 {
		biaya *= orang
	}
	return biaya
}

func main() {
	var M int
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&M)

	for i := 1; i <= M; i++ {
		var jummenu, jumorang, jumsisa int
		fmt.Printf("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya):")
		fmt.Scan(&jummenu, &jumorang, &jumsisa)
		biaya := totalBiaya(jummenu, jumorang, jumsisa)
		fmt.Printf("Total biaya untuk rombongan %d: Rp %d\n", i, biaya)
	}
}