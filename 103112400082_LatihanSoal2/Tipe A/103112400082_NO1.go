// Rizkina Azizah_103112400082
package main

import "fmt"

func hitungBiayaSewa(jam int, menit int, member bool, voucher string) float64 {
	var harga float64
	if member {
		harga = 3500.0 
	} else {
		harga = 5000.0 
	}

	totalDurasiJam := float64(jam) + float64(menit)/60.0

	biayaSewa := totalDurasiJam * harga

	//diskon
	if totalDurasiJam >= 3 && (len(voucher) == 5 || len(voucher) == 6) {
		biayaSewa *= 0.9 
	}

	return biayaSewa
}

func main() {
	var jam, menit int
	var member bool
	var voucher string

	
	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&jam)
	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&menit)
	fmt.Print("Apakah member (true/false): ")
	fmt.Scan(&member)
	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scan(&voucher)

	biaya := hitungBiayaSewa(jam, menit, member, voucher)

	fmt.Printf("Biaya sewa: Rp %.2f\n", biaya)
}