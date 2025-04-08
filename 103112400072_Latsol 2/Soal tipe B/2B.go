package main

import "fmt"

func main() {
	var jumlahRombongan int
	
	
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&jumlahRombongan)
	
	
	for i := 1; i <= jumlahRombongan; i++ {
		var jumlahMenu, jumlahOrang int
		var sisaMakanan int
		
		fmt.Printf("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk ya)\n")
		fmt.Scan(&jumlahMenu, &jumlahOrang, &sisaMakanan)
		fmt.Printf(": %d %d %d\n", jumlahMenu, jumlahOrang, sisaMakanan)
		
		
		totalBiaya := hitungBiaya(jumlahMenu, jumlahOrang, sisaMakanan)
		
		
		fmt.Printf("Total biaya untuk rombongan %d: Rp %d\n", i, totalBiaya)
	}
}

func hitungBiaya(jumlahMenu, jumlahOrang, sisaMakanan int) int {
	var biaya int
	
	
	if jumlahMenu <= 3 {
		biaya = 10000
	} else if jumlahMenu <= 50 {
		biaya = 10000 + (jumlahMenu - 3) * 2500
	} else {
		biaya = 100000
	}
	
	
	if sisaMakanan == 1 {
		biaya = biaya * jumlahOrang
	}
	
	return biaya
}