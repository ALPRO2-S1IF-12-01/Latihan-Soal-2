// Savila Nur Fadilla
// 103112400031

package main

import "fmt"

func main() {
	var angka int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&angka)

	kiri, kanan := potongAngka(angka)
	fmt.Println("Bilangan 1:", kiri)
	fmt.Println("Bilangan 2:", kanan)
	fmt.Print("Hasil penjumlahan:", kiri+kanan)
}

func potongAngka(angka int) (int, int) {
	panjang := hitungDigit(angka)
	posisi := panjang / 2
	if panjang % 2 != 0 {
		posisi++
	}

	pembagi := 1
	for i := 0; i < panjang-posisi; i++ {
		pembagi *= 10
	}

	kiri := angka / pembagi
	kanan := angka % pembagi

	return kiri, kanan
}

func hitungDigit(n int) int {
	hitung := 0
	for n > 0 {
		n /= 10
		hitung++
	}
	return hitung
}
