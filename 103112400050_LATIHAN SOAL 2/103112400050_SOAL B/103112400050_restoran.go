package main

//103112400050
import "fmt"

func main() {
	var x int
	fmt.Print("masukkan jumlah rombongan: ")
	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		var menu, orang, sisa int
		fmt.Printf("masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya)\n: ")
		fmt.Scan(&menu, &orang, &sisa)
		var tarif int
		switch {
		case menu > 50:
			tarif = 100000
		case menu <= 3:
			tarif = 10000
		default:
			tarif = 10000 + (menu-3)*2500
		}
		if sisa == 1 {
			tarif *= orang
		}
		fmt.Printf("total biaya untuk rombongan %d: Rp %d\n", i, tarif)
	}
}
