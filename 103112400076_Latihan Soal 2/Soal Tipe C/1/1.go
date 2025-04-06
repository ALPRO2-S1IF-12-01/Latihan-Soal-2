package main

import (
	"fmt"
	"math"
)

func main() {
	NIM := "103112400076"
	fmt.Println("NIM:", NIM)
	var input int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&input)
	jumlahDigit := 0
	temp := input
	for temp > 0 {
		jumlahDigit++
		temp /= 10
	}
	var kananDigit int
	if jumlahDigit%2 == 0 {
		kananDigit = jumlahDigit / 2
	} else {
		kananDigit = jumlahDigit / 2
	}
	pangkat := int(math.Pow(10, float64(kananDigit)))
	kiri := input / pangkat
	kanan := input % pangkat
	fmt.Println("Bilangan 1:", kiri)
	fmt.Println("Bilangan 2:", kanan)
	fmt.Println("Hasil penjumlahan:", kiri+kanan)
}
