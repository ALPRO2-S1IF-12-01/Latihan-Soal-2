package main

import "fmt"

func hitungBiayaSewa(jam int, menit int, member bool, kodeVoucher string) float64 {
	totalMenit := jam*60 + menit
	durasiJam := float64(totalMenit) / 60.0

	var biayaPerJam float64
	if member {
		biayaPerJam = 3500
	} else {
		biayaPerJam = 5000
	}

	biayaSewa := durasiJam * biayaPerJam
	if durasiJam > 3 {
		lastDigit := 0
		if len(kodeVoucher) > 0 {
			lastDigit = int(kodeVoucher[len(kodeVoucher)-1] - '0')
		}
		if lastDigit == 5 || lastDigit == 6 {
			biayaSewa *= 0.9
		}
	}

	return biayaSewa
}

func main() {
	var jam, menit int
	var isMemberStr string
	var isMember_13112400079 bool
	var kodeVoucher string

	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&jam)

	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&menit)

	fmt.Print("Apakah member? (true/false): ")
	fmt.Scan(&isMemberStr)
	if isMemberStr == "true" {
		isMember_13112400079 = true
	} else if isMemberStr == "false" {
		isMember_13112400079 = false
	}

	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scan(&kodeVoucher)

	biayaSetelahDiskon := hitungBiayaSewa(jam, menit, isMember_13112400079, kodeVoucher)

	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", biayaSetelahDiskon)
}
