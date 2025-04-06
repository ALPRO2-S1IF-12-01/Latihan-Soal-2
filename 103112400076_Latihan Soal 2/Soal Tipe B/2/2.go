package main

import (
	"fmt"
)
func main() {
	var m int
	var NIM string = "103112400076" 
	fmt.Println("NIM :", NIM)
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&m)
	for i := 1; i <= m; i++ {
		var menu, orang, sisa int
		fmt.Printf("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya):\n: ")
		fmt.Scan(&menu, &orang, &sisa)
		tarif := 0
		if menu <= 3 {
			tarif = 10000
		} else if menu <= 50 {
			tarif = 10000 + (menu-3)*2500
		} else {
			tarif = 100000
		}
		if sisa == 1 {
			tarif *= orang
		}
		fmt.Printf("Total biaya untuk rombongan %d: Rp %d\n", i, tarif)
	}
}
