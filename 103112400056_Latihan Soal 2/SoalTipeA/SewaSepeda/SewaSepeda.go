// Felix Pedrosa V

package main

import "fmt"

// Subprogram untuk menghitung biaya sewa
func hitungBiayaSewa(jam int, menit int, member bool, noVoucher string) float64 {
	// Konstanta tarif
	var tarif float64
	if member {
		tarif = 3500
	} else {
		tarif = 5000
	}

	// Hitung total durasi dalam jam
	totalMeni := jam*60 + menit
	var totalJam int
	if totalMeni%60 >= 10 {
		totalJam = totalMeni/60 + 1
	} else {
		totalJam = totalMeni / 60
	}

	// Hitung biaya sebelum diskon
	biaya := float64(totalJam) * tarif

	// Cek memenuhi syarat diskon
	if len(noVoucher) >= 5 && len(noVoucher) <= 6 {
		diskon := 0.1 * biaya
		biaya -= diskon
	}

	return biaya
}

func main() {
	// Input
	var jam, menit int
	var member_103112400056 bool
	var noVoucher string

	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&jam)
	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&menit)
	fmt.Print("Apakah member? (true/false): ")
	fmt.Scan(&member_103112400056)
	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scan(&noVoucher)

	// Hitung biaya sewa
	biaya := hitungBiayaSewa(jam, menit, member_103112400056, noVoucher)

	// Output
	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", biaya)
}