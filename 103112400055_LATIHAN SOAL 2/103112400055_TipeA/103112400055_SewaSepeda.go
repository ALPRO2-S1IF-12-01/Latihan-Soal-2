package main

import "fmt"

func hitungBiayaSewa(jam int, menit int, isMember bool, nomorVoucher string) float64 {
	const tarifNonMember = 5000.0
	const tarifMember = 3500.0
	const diskonPersen = 0.10

	totalMenit := jam*60 + menit
	if totalMenit < 60 {
		return 0.0
	}

	var tarif float64
	if isMember {
		tarif = tarifMember
	} else {
		tarif = tarifNonMember
	}

	durasiSewa := totalMenit / 60
	if totalMenit%60 > 0 {
		durasiSewa++
	}

	biayaSewa := float64(durasiSewa) * tarif

	if len(nomorVoucher) >= 5 && len(nomorVoucher) <= 6 {
		biayaSewa = biayaSewa * (1 - diskonPersen)
	}

	return biayaSewa
}

func main() {
	var jam_103112400055, menit int
	var isMember bool
	var nomorVoucher string

	// Input pengguna
	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&jam_103112400055)
	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&menit)
	fmt.Print("Apakah member? (true/false): ")
	fmt.Scan(&isMember)

	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scan(&nomorVoucher)

	biaya := hitungBiayaSewa(jam_103112400055, menit, isMember, nomorVoucher)
	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", biaya)
}
