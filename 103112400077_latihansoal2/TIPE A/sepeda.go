package main

import "fmt"

func tarif(jam, menit int, member bool, voucher int) float64{
	var tarifJam float64
	if member {
		tarifJam = 3500
	} else {
		tarifJam = 5000
	}

	totalJam := float64(jam)
	if jam == 0 && menit > 0 {
		totalJam = 1
	} else if menit >= 10 {
		totalJam += float64(menit) / 60 
	}

	totalTarif := totalJam * tarifJam

	if (voucher >= 10000 && voucher <= 99999) ||  (voucher >= 100000 && voucher <= 999999) {
		totalTarif *= 0.9
	}

	return totalTarif
}

func main() {
	var jam_103112400077, menit, voucher int
	var member bool

	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&jam_103112400077)
	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&menit)
	fmt.Print("Apakah member? (true/false): ")
	fmt.Scan(&member)
	fmt.Print("Masukkan nomor voucher (jika ada): ")
	fmt.Scan(&voucher)

	biaya := tarif(jam_103112400077, menit, member, voucher)

	fmt.Printf("Biaya sewa setelah diskon (jika memenuhi syarat): Rp %.2f\n", biaya)

}