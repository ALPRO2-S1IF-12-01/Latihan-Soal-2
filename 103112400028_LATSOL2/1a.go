// MUAHAMMAD GAMEL AL GHIFARI
// 103112400028
package main

import (
	"fmt"
)

func main() {
	var jam, menit int
	var member bool
	var voucher string

	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&jam)
	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&menit)
	fmt.Print("Apakah member? (true/false): ")
	fmt.Scan(&member)
	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scan(&voucher)

	totalJam := hitungTotalJam(jam, menit)

	biaya := hitungBiaya(totalJam, member)

	if totalJam > 3 {
		biaya = biaya * 0.9
	}

	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", biaya)
}

func hitungTotalJam(jam, menit int) float64 {
	totalJam := float64(jam)
	if menit >= 10 || totalJam < 1 {
		totalJam += float64(menit) / 60.0
	}
	return totalJam
}

func hitungBiaya(totalJam float64, member bool) float64 {
	tarif := 5000.0
	if member {
		tarif = 3500.0
	}
	return totalJam * tarif
}
