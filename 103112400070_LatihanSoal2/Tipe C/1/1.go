// Achmad Zulvan Nur Hakim 103112400070
package main

import "fmt"

func potongDanJumlah(angka int) (int, int, int) {

	depan := angka
	jmldigit := 0
	for belakang := depan; belakang > 0; belakang /= 10 {
		jmldigit++
	}

	ambilDepan := (jmldigit + 1) / 2


	pangkat := 1
	for i := 0; i < jmldigit-ambilDepan; i++ {
		pangkat *= 10
	}

	bagianKiri := angka / pangkat
	bagianKanan := angka % pangkat
	total := bagianKiri + bagianKanan

	return bagianKiri, bagianKanan, total
}

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&bilangan)

	angka1, angka2, jumlah := potongDanJumlah(bilangan)

	fmt.Printf("Bilangan 1: %d\n", angka1)
	fmt.Printf("Bilangan 2: %d\n", angka2)
	fmt.Printf("Hasil penjumlahan: %d\n", jumlah)
}
