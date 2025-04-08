package main

import "fmt"

func main() {
	fmt.Println("Nama: Raja Muhammad Lufhti\nNim : 103112400027")
	var angka int
	fmt.Print("Masukkan bilangan bulat positif (>10): ")
	fmt.Scan(&angka)

	bagianKiri, bagianKanan := potongAngka(angka)
	hasilJumlah := bagianKiri + bagianKanan

	fmt.Println("Bilangan 1:", bagianKiri)
	fmt.Println("Bilangan 2:", bagianKanan)
	fmt.Println("Hasil penjumlahan:", hasilJumlah)
}

func potongAngka(angka int) (int, int) {
	panjang := 0
	salinan := angka
	for salinan > 0 {
		salinan /= 10
		panjang++
	}

	indeksPotong := (panjang + 1) / 2 

	salinan = angka
	faktor := 1
	for i := 0; i < panjang-indeksPotong; i++ {
		faktor *= 10
	}

	bagianKiri := angka / faktor
	bagianKanan := angka % faktor

	return bagianKiri, bagianKanan
}
