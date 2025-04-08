package main

import "fmt"

func tarif(menu, orang, sisaMakanan int) int {
	var tarifRombongan, tarifMenu, tarif int
	
	tarifRombongan = 10000
	if menu > 3 {
		tarifMenu = (menu-3) * 2500
	}
	if menu > 50 {
		tarifMenu = 100000

	}
	if sisaMakanan == 1 {
		tarif = (tarifRombongan+tarifMenu)*orang
	} else {
		tarif = (tarifRombongan+tarifMenu)
	}
	return tarif
}
func main() {
	var M_103112400041, menu, orang, sisaMakanan int
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&M_103112400041)
	rombongan:=0
	for i := 1; i <= M_103112400041; i++ {
		fmt.Print("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya): ")
		fmt.Scan(&menu, &orang, &sisaMakanan)
		rombongan=rombongan+1
		fmt.Printf("Total biaya untuk rombongan %d: Rp %d", rombongan, tarif(menu, orang, sisaMakanan))
		fmt.Println() 
	}
}