/*
Nama: Dimas Ramadhani
NIM: 103112400065
*/

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
	var hargaPerjam, harga float64
	if member == true {
		hargaPerjam = 3500.0
	} else {
		hargaPerjam = 5000.0
	}
	harga = durasi * hargaPerjam
	if durasi > 3 && cekVoucher(voucher) {
		harga = harga * 9 / 10
	}
	return harga
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
	var durasi float64
	if jam < 1 {
		durasi = 1
	} else if jam >= 1 && menit <= 10 {
		durasi = float64(jam)
	} else if jam >= 1 && menit > 10 {
		durasi = float64(jam) + float64(menit)/60
	}
	return durasi
}
