// Nama: Anggun Wahyu Widiyana
// NIM : 103112480280
package main

import (
	"fmt"
)

func main() {
	var bilangan_103112480280 int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&bilangan_103112480280)

	strBil := fmt.Sprintf("%d", bilangan_103112480280)
	panjang := len(strBil)

	var bagi int
	if panjang%2 == 0 {
		bagi = panjang / 2
	} else {
		bagi = (panjang / 2) + 1
	}

	bil1Str := strBil[:bagi]
	bil2Str := strBil[bagi:]

	var bil1, bil2 int
	fmt.Sscanf(bil1Str, "%d", &bil1)
	fmt.Sscanf(bil2Str, "%d", &bil2)

	fmt.Printf("Bilangan 1: %d\n", bil1)
	fmt.Printf("Bilangan 2: %d\n", bil2)
	fmt.Printf("Hasil penjumlahan: %d\n", bil1+bil2)
}