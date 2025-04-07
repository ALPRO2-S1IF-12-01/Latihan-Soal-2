// JESIKA MMETANIA RAHMA ARIFIN
// 103112400080

// JESIKA METANIA RAHMA ARIFIN
// 103112400080

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var angka int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&angka)

	stringAngka := strconv.Itoa(angka)

	if len(stringAngka) <= 1 {
		fmt.Println("Bilangan harus lebih dari 10.")
		return
	}
	tengah := (len(stringAngka) + 1) / 2

	
	bagianDepan := stringAngka[:tengah]
	bagianBelakang := stringAngka[tengah:]

	angkaDepan, _ := strconv.Atoi(bagianDepan)
	angkaBelakang, _ := strconv.Atoi(bagianBelakang)

	
	fmt.Println("Angka bagian depan:", angkaDepan)
	fmt.Println("Angka bagian belakang:", angkaBelakang)
	fmt.Println("Hasil penjumlahan:", angkaDepan+angkaBelakang)
}
