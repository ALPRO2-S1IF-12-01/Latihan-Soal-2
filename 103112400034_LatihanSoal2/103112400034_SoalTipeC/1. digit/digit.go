package main

import (
	"fmt"
	"math"
)

func main() {
	var angka int
	fmt.Println("NAMA: MULIA AKBAR NANDA PRATAMA\nNIM: 103112400034")
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&angka)

	jumlahDigit := 0
	salinan := angka
	for salinan > 0 {
		salinan /= 10
		jumlahDigit++
	}

	tengah := jumlahDigit / 2
	if jumlahDigit%2 != 0 {
		tengah++ 
	}
	pembagi := int(math.Pow(10, float64(jumlahDigit-tengah)))
	bilangan1 := angka / pembagi
	bilangan2 := angka % pembagi
	fmt.Println("Bilangan 1:", bilangan1)
	fmt.Println("Bilangan 2:", bilangan2)
	fmt.Println("Hasil penjumlahan:", bilangan1+bilangan2)
}
