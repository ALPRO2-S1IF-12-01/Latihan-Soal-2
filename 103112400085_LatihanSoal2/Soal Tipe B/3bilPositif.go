//Anastasia Adinda N.I
package main

import (
	"fmt"
)

func KER_103112400085(data []int, posisi int, jumlah int) int {
	if posisi >= len(data) {
		return jumlah
	}
	if data[posisi] > 0 && data[posisi]%4 == 0 {
		jumlah += data[posisi]
	}
	return KER_103112400085(data, posisi+1, jumlah)
}

func main() {
	var angka int
	var listAngka [100]int
	var jumlahData int

	fmt.Println("Masukkan bilangan (negatif untuk berhenti):")

	for {
		fmt.Scan(&angka)
		if angka < 0 || jumlahData >= 100 {
			break
		}
		listAngka[jumlahData] = angka
		jumlahData++
	}

	total := KER_103112400085(listAngka[:jumlahData], 0, 0)
	fmt.Printf("Jumlah bilangan kelipatan 4: %d\n", total)
}
