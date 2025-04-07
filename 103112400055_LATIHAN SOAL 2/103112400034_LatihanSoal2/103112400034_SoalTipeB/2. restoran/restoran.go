package main

import "fmt"

func main() {
	var jr int
	fmt.Println("NAMA: MULIA AKBAR NANDA PRATAMA\nNIM: 103112400034")
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scan(&jr)

	for i := 1; i <= jr; i++ {
		var jm, jo, sisa int
		fmt.Printf("Masukkan jumlah menu, jumlah orang, dan status sisa makanan (0 untuk tidak, 1 untuk iya) \nuntuk rombongan %d: ", i)
		fmt.Scan(&jm, &jo, &sisa)

		totalTarif := 0

		if jm > 50 {
			totalTarif = 100000
		} else if jm > 3 {
			totalTarif = 10000 + (jm-3)*2500
		} else {
			totalTarif = 10000
		}

		if sisa == 1 {
			totalTarif *= jo
		}

		fmt.Printf("Total biaya untuk rombongan %d: Rp %d\n", i, totalTarif)
	}
}
