// Savila Nur Fadilla
// 103112400031

package main

import "fmt"

func main() {
	var jam, menit int
	var voucher string
	var member bool

	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&jam)
	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&menit)
	fmt.Print("Apakah member? (true/false): ")
	fmt.Scan(&member)
	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scan(&voucher)

	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f", hitungBiaya(jam, menit, member, voucher))
}

func hitungBiaya(jam, menit int, member bool, voucher string) float64 {
	durasi := totalJam(jam, menit)

	perJam := 5000.00
	if member {
		perJam = 3500.00
	}

	biaya := durasi * perJam

	if durasi > 3 && cekVoucher(voucher) {
		biaya *= 0.9
	}

	return biaya
}

func cekVoucher(voucher string) bool {
	var hasil bool
	digit := len(voucher)
	if digit >= 5 {
		hasil = true
	} else {
		hasil = false
	}
	return hasil
}

func totalJam(jam, menit int) float64 {
	totalMenit := (jam * 60) + menit
	durasi := float64(totalMenit) / 60

	if durasi < 1 {
		return 1
	}
	return durasi
}
