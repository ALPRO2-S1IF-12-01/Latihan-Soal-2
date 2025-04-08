package main

import "fmt"

func main() {
	var (
		menu, orang, rombongan_103112400061, sisa int
	)
	fmt.Print("Masukkan jumlah rombongan_103112400061: ")
	fmt.Scan(&rombongan_103112400061)
	for i := 1; i <= rombongan_103112400061; i++ {
		fmt.Print("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya): ")
		fmt.Scan(&menu, &orang, &sisa)
		fmt.Printf("Total biaya untuk rombongan %v: Rp %v\n", i, restoran(menu, orang, sisa))
	}
}

func restoran(menu, orang, sisa int) int {
	var biaya int
	if menu <= 3 && sisa == 0 {
		biaya = 10000
	} else if menu > 3 && sisa == 0 {
		biaya = 10000
		for i := menu; i > 3; i-- {
			biaya += 2500
		}
	}
	
	if menu <= 3 && sisa == 1 {
		biaya = 10000 * orang
	} else if menu > 3 && sisa == 1 {
		biaya = 10000
		for i := menu; i > 3; i-- {
			biaya += 2500
		}
		biaya *= orang
	}
	
	return biaya
}