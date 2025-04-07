package main

import (
	"fmt"
)

func main() {
	fmt.Println("NAMA: MULIA AKBAR NANDA PRATAMA\nNIM: 103112400034")
	fmt.Print("Masukkan angka (negatif untuk berhenti):\n ")
	total := hasil(0)
	fmt.Println("Jumlah bilangan kelipatan 4:", total)
}

func hasil(total int) int {
	var angka int
	fmt.Scan(&angka)
	if angka < 0 {
		return total
	}
	if angka > 0 && angka%4 == 0 {
		total +=angka
	}
	return hasil(total)
}

