package main

import "fmt"

func main() {
	var x, y int
	fmt.Println("NAMA: MULIA AKBAR NANDA PRATAMA\nNIM: 103112400034")

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	jumlahPertemuan := 0
	for hari := 1; hari <= 365; hari++ {
		if hari%x == 0 && hari%y != 0 {
			jumlahPertemuan++
		}
	}
	fmt.Println("Jumlah pertemuan dalam setahun:", jumlahPertemuan)
}
