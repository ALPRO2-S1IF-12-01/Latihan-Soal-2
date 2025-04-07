// RYAN AKEYLA NOVIANTO WIDODO
// 103112400081

package main

import (
	"fmt"
	"math"
)

func hitungBiayaSewa(jam, menit int, member bool, noVoucher string) float64 {
	var tarif float64
	if member {
		tarif = 3500
	} else {
		tarif = 5000
	}

	totalJam := math.Ceil(float64(jam*60+menit) / 60)
	biaya := totalJam * tarif

	if len(noVoucher) >= 5 && len(noVoucher) <= 6 {
		biaya -= 0.1 * biaya
	}

	return biaya
}

func main() {
	var jam, menit int
	var member bool
	var noVoucher string

	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&jam)
	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&menit)
	fmt.Print("Apakah member? (true/false): ")
	fmt.Scan(&member)
	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scan(&noVoucher)

	biaya := hitungBiayaSewa(jam, menit, member, noVoucher)
	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", biaya)
}
