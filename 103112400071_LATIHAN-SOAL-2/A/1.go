// Raihan Adi Arba_103112400071
package main

import "fmt"

// Fungsi untuk menghitung biaya sewa
func hitungBiayaSewa(DurasiJam int, DurasiMenit int, INIMember bool, nomorVoucher string) float64 {
	tarifPerJam := 5000.0
	if INIMember {
		tarifPerJam = 3500.0
	}

	// Total jam sebagai float
	totalJam := float64(DurasiJam) + float64(DurasiMenit)/60.0

	// Minimal durasi adalah 1 jam
	if totalJam < 1.0 {
		totalJam = 1.0
	}

	biayaTotal := totalJam * tarifPerJam

	// Diskon hanya jika durasi > 3 jam dan voucher valid
	if totalJam > 3.0 && isVoucherValid(nomorVoucher) {
		biayaTotal *= 0.9
	}

	return biayaTotal
}

// Fungsi validasi voucher (5 atau 6 digit angka)
func isVoucherValid(voucher string) bool {
	jumlahDigit := 0
	for _, ch := range voucher {
		if ch < '0' || ch > '9' {
			return false
		}
		jumlahDigit++
	}
	return jumlahDigit == 5 || jumlahDigit == 6
}

func main() {
	var DurasiJam, DurasiMenit int
	var INIMember bool
	var nomorVoucher string

	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&DurasiJam)

	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&DurasiMenit)

	fmt.Print("Apakah member? (true/false): ")
	fmt.Scan(&INIMember)

	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scan(&nomorVoucher)

	biayaSewa := hitungBiayaSewa(DurasiJam, DurasiMenit, INIMember, nomorVoucher)
	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", biayaSewa)
}
