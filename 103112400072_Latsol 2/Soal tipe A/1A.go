package main

import "fmt"

func main() {
	var durasiJam, durasiMenit int
	var kodeVoucher string
	var isMember bool

	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&durasiJam)

	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&durasiMenit)

	fmt.Print("Apakah member? (true/false): ")
	fmt.Scan(&isMember)

	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scan(&kodeVoucher)

	totalBayar := hitungBiayaSewa(durasiJam, durasiMenit, isMember, kodeVoucher)
	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", totalBayar)
}

func hitungBiayaSewa(jam, menit int, member bool, voucher string) float64 {
	totalDurasi := hitungTotalDurasi(jam, menit)

	tarifPerJam := 5000.00
	if member {
		tarifPerJam = 3500.00
	}

	totalBiaya := float64(totalDurasi) * tarifPerJam

	if totalDurasi >= 3 && panjangVoucherValid(voucher) {
		totalBiaya *= 0.9
	}

	return totalBiaya
}

func panjangVoucherValid(voucher string) bool {
	return len(voucher) >= 5
}

func hitungTotalDurasi(jam, menit int) int {
	if jam == 0 && menit < 10 {
		return 1
	} else if menit >= 10 {
		return jam + 1
	}
	return jam
}
