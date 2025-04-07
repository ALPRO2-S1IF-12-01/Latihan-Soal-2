// M.HANIF AL FAIZ
// 103112400042
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&bilangan)

	// Konversi bilangan ke string untuk memudahkan pemotongan
	strBilangan := strconv.Itoa(bilangan)
	panjang := len(strBilangan)

	// Tentukan titik pemotongan
	titikPotong := panjang / 2
	if panjang%2 != 0 {
		titikPotong = (panjang / 2) + 1
	}

	// Potong bilangan menjadi dua bagian
	bagian1 := strBilangan[:titikPotong]
	bagian2 := strBilangan[titikPotong:]

	// Konversi kembali ke integer
	bil1, _ := strconv.Atoi(bagian1)
	bil2, _ := strconv.Atoi(bagian2)

	// Hitung penjumlahan
	hasilJumlah := bil1 + bil2

	// Tampilkan output
	fmt.Printf("Bilangan 1: %s\n", bagian1)
	fmt.Printf("Bilangan 2: %s\n", bagian2)
	fmt.Printf("Hasil penjumlahan: %d\n", hasilJumlah)
}
